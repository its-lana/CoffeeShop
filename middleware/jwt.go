package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/its-lana/coffee-shop/apperr"
	"github.com/its-lana/coffee-shop/common"
	"github.com/its-lana/coffee-shop/dto"
)

func AuthorizeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if gin.Mode() == gin.TestMode {
			return
		}

		var response dto.Response

		header := ctx.Request.Header["Authorization"]
		if len(header) == 0 {
			response.Message = apperr.ErrBearerTokenInvalid.Error()
			ctx.AbortWithStatusJSON(apperr.ErrBearerTokenInvalid.Code, response)
			return
		}

		splittedHeader := strings.Split(header[0], " ")
		if len(splittedHeader) != 2 {
			response.Message = apperr.ErrUnathorized.Error()
			ctx.AbortWithStatusJSON(apperr.ErrUnathorized.Code, response)
			return
		}

		claims := &dto.JWTClaims{}

		token, err := jwt.ParseWithClaims(splittedHeader[1], claims, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, apperr.ErrWrongCredentials
			}
			return []byte(os.Getenv("API_SECRET")), nil
		})
		if err != nil {
			response.Message = err.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		_, ok := token.Claims.(*dto.JWTClaims)
		if !ok || !token.Valid {
			response.Message = apperr.ErrUnathorized.Error()
			ctx.AbortWithStatusJSON(apperr.ErrUnathorized.Code, response)
			return
		}

		newCtx := context.WithValue(ctx.Request.Context(), common.ID, claims.ID)
		newCtx = context.WithValue(newCtx, common.Role, claims.Role)
		ctx.Request = ctx.Request.WithContext(newCtx)
		ctx.Next()
	}
}
