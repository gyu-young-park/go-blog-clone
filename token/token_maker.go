package token

import "time"

type TokenMaker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	ValidateToken(token string) (*Claim, error)
}
