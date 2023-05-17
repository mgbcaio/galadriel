// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3"
)

// Defines values for ConsentStatus.
const (
	Accepted ConsentStatus = "accepted"
	Denied   ConsentStatus = "denied"
	Pending  ConsentStatus = "pending"
)

// ApiError defines model for ApiError.
type ApiError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// BundleDigest base64 encoded SHA-256 digest of the bundle
type BundleDigest = string

// Certificate X.509 certificate in PEM format
type Certificate = string

// ConsentStatus defines model for ConsentStatus.
type ConsentStatus string

// JWT defines model for JWT.
type JWT = string

// JoinToken defines model for JoinToken.
type JoinToken = UUID

// Relationship defines model for Relationship.
type Relationship struct {
	CreatedAt           time.Time        `json:"created_at"`
	Id                  UUID             `json:"id"`
	TrustDomainAConsent ConsentStatus    `json:"trust_domain_a_consent"`
	TrustDomainAId      UUID             `json:"trust_domain_a_id"`
	TrustDomainAName    *TrustDomainName `json:"trust_domain_a_name,omitempty"`
	TrustDomainBConsent ConsentStatus    `json:"trust_domain_b_consent"`
	TrustDomainBId      UUID             `json:"trust_domain_b_id"`
	TrustDomainBName    *TrustDomainName `json:"trust_domain_b_name,omitempty"`
	UpdatedAt           time.Time        `json:"updated_at"`
}

// SPIFFEID defines model for SPIFFEID.
type SPIFFEID = string

// Signature defines model for Signature.
type Signature = string

// TrustBundle X.509 certificate in PEM format
type TrustBundle = Certificate

// TrustDomain defines model for TrustDomain.
type TrustDomain struct {
	CreatedAt         time.Time       `json:"created_at"`
	Description       *string         `json:"description,omitempty"`
	HarvesterSpiffeId *SPIFFEID       `json:"harvester_spiffe_id,omitempty"`
	Id                UUID            `json:"id"`
	Name              TrustDomainName `json:"name"`

	// OnboardingBundle X.509 certificate in PEM format
	OnboardingBundle *TrustBundle `json:"onboarding_bundle,omitempty"`
	UpdatedAt        time.Time    `json:"updated_at"`
}

// TrustDomainName defines model for TrustDomainName.
type TrustDomainName = string

// UUID defines model for UUID.
type UUID = openapi_types.UUID

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
