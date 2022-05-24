package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const MIN_SECRET_KET_SIZE = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (TokenMaker, error) {
	if len(secretKey) < MIN_SECRET_KET_SIZE {
		return nil, fmt.Errorf("Invalid to get secretKey size[%v]\n", len(secretKey))
	}
	jwtMaker := &JWTMaker{
		secretKey: secretKey,
	}
	return jwtMaker, nil
}

func (jwtMaker *JWTMaker) CreateToken(email string, duration time.Duration) (string, error) {
	claim := NewClaim(email, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return jwtToken.SignedString([]byte(jwtMaker.secretKey))
}

func (jwtMaker *JWTMaker) ValidateToken(token string) (*Claim, error) {
	// check algo
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(jwtMaker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Claim{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	claim, ok := jwtToken.Claims.(*Claim)
	if !ok {
		return nil, ErrInvalidToken
	}
	return claim, nil
}
