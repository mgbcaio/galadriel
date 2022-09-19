package util

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/HewlettPackard/galadriel/pkg/common"
	"github.com/HewlettPackard/galadriel/pkg/server/datastore"
	"github.com/google/uuid"
	"io"
	"net"
	"net/http"
)

// ServerLocalClient represents a local client of the Galadriel Server.
type ServerLocalClient interface {
	CreateMember(m common.Member) (*datastore.Member, error)
	CreateRelationship(r common.Relationship) (*datastore.Relationship, error)
	GenerateAccessToken(memberID uuid.UUID) (*datastore.AccessToken, error)
}

// TODO: improve this adding options for the transport, dialcontext, and http.Client.
func NewServerClient(socketPath string) ServerLocalClient {
	t := &http.Transport{
		DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
			var d net.Dialer
			return d.DialContext(ctx, "unix", socketPath)
		}}
	c := &http.Client{
		Transport: t,
	}

	return serverClient{client: c}

}

type serverClient struct {
	client *http.Client
}

func (c serverClient) CreateMember(m common.Member) (*datastore.Member, error) {
	err := validateMember(m)
	if err != nil {
		return nil, err
	}

	memberBytes, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	bytes.NewBuffer(memberBytes)
	r, err := c.client.Post("http://unix/createMember", "application/json", bytes.NewBuffer(memberBytes))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	createdMember := &datastore.Member{}
	err = json.Unmarshal(b, createdMember)
	if err != nil {
		return nil, err
	}

	return createdMember, nil
}

func (c serverClient) CreateRelationship(rel common.Relationship) (*datastore.Relationship, error) {
	err := validateRelationship(rel)
	if err != nil {
		return nil, err
	}

	relBytes, err := json.Marshal(rel)
	if err != nil {
		return nil, err
	}

	bytes.NewBuffer(relBytes)
	r, err := c.client.Post("http://unix/createRelationship", "application/json", bytes.NewBuffer(relBytes))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	createdRelationship := &datastore.Relationship{}
	err = json.Unmarshal(b, createdRelationship)
	if err != nil {
		return nil, err
	}

	return createdRelationship, nil
}

func (c serverClient) GenerateAccessToken(memberID uuid.UUID) (*datastore.AccessToken, error) {
	b, err := json.Marshal(common.Member{ID: memberID})
	if err != nil {
		return nil, err
	}
	r, err := c.client.Post("http://unix/token", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	createdToken := &datastore.AccessToken{}
	err = json.Unmarshal(body, createdToken)
	if err != nil {
		return nil, err
	}
	return createdToken, nil
}

func validateMember(m common.Member) error {
	// TODO: checks
	return nil
}

func validateRelationship(m common.Relationship) error {
	// TODO: checks
	return nil
}
