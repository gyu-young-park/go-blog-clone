package token

import (
	"errors"
	"time"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token has invalid")
)

type Claim struct {
	Email     string    `json:"email"`
	ExpiredAt time.Time `json:"expired_at"`
	IssuedAt  time.Time `json:"issued_at"`
}

func NewClaim(email string, duration time.Duration) *Claim {
	claim := &Claim{
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return claim
}

func (claim *Claim) Valid() error {
	if time.Now().After(claim.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
