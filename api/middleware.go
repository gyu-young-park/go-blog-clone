package api

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gyu-young-park/go_blog/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationClaimKey   = "authorization_claim"
)

func authMiddleware(tokenMaker token.TokenMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeaderKey := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeaderKey) == 0 {
			err := errors.New("authorizationHeader is not provided")
			ctx.AbortWithStatusJSON(401, errResponse(err))
			return
		}
		fields := strings.Fields(authorizationHeaderKey)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(401, errResponse(err))
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(401, errResponse(err))
			return
		}
		accessToken := fields[1]
		log.Println(accessToken)
		claim, err := tokenMaker.ValidateToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(401, errResponse(err))
			return
		}
		ctx.Set(authorizationClaimKey, claim)
		ctx.Next()
	}
}
