package telemetry

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	prommetrics "github.com/armon/go-metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

type prometheusRunner struct {
	c      *PrometheusConfig
	log    logrus.FieldLogger
	server *http.Server
	sink   Sink
}

func newPrometheusRunner(c *MetricsConfig) (sinkRunner, error) {
	runner := &prometheusRunner{
		c:   c.PrometheusConf,
		log: c.Logger,
	}

	if runner.c == nil {
		return nil, errors.New("configuration file not specified")
	}

	var err error
	runner.sink, err = prommetrics.NewPrometheusSinkFrom(prommetrics.PrometheusOpts{})
	if err != nil {
		return nil, err
	}

	handlerOpts := promhttp.HandlerOpts{
		ErrorLog: runner.log,
	}
	handler := promhttp.HandlerFor(prometheus.DefaultGatherer, handlerOpts)

	if runner.c.Host == "" {
		runner.c.Host = "localhost"
	}

	runner.server = &http.Server{
		Addr:              fmt.Sprintf("%s:%d", runner.c.Host, runner.c.Port),
		Handler:           handler,
		ReadHeaderTimeout: time.Second * 10,
	}

	return runner, nil
}

func (p *prometheusRunner) isConfigured() bool {
	return p.c != nil
}

func (p *prometheusRunner) sinks() []Sink {
	if !p.isConfigured() {
		return []Sink{}
	}

	return []Sink{p.sink}
}

func (p *prometheusRunner) run(ctx context.Context) error {
	if !p.isConfigured() {
		return nil
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := p.server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			p.log.Warnf("Prometheus listener stopped unexpectedly: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		p.server.Close()
	}()

	wg.Wait()
	return ctx.Err()
}
