package middleware

import (
	"net/http"
	"project-tiga/helper"
	"project-tiga/model"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")

	token := strings.Split(auth, " ")[1]

	jwtToken, err := helper.VerifyAccessToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	_, ok = claims["refresh"]
	if ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorInvalidToken)
		return
	}

	ctx.Set("user_id", claims["user_id"])

	ctx.Next()
}

func AuthRefreshMiddleware(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")

	token := strings.Split(auth, " ")[1]

	jwtToken, err := helper.VerifyRefreshToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	v, ok := claims["refresh"]
	if !ok || !v.(bool) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorInvalidToken)
		return
	}

	ctx.Set("user_id", claims["user_id"])

	ctx.Next()
}
