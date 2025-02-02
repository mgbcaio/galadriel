// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package sqlite

import (
	"database/sql"
	"time"
)

type Bundle struct {
	ID                 string
	TrustDomainID      string
	Data               []byte
	Digest             []byte
	Signature          []byte
	SigningCertificate []byte
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type JoinToken struct {
	ID            string
	TrustDomainID string
	Token         string
	Used          bool
	ExpiresAt     time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Relationship struct {
	ID                  string
	TrustDomainAID      string
	TrustDomainBID      string
	TrustDomainAConsent string
	TrustDomainBConsent string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type TrustDomain struct {
	ID          string
	Name        string
	Description sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
