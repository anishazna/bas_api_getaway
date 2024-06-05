package handler

import (
	"api_getaway/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authImplement struct{}

type authInterface interface {
	Login(*gin.Context)
}

func NewAuthRepository() authInterface {
	return &authImplement{}
}

type BodyPayloadAuth struct {
	username string
	password string
}

func (a *authImplement) Login(g *gin.Context) {
	bodyPayloadAuth := BodyPayloadAuth{}

	err := g.BindJSON(&bodyPayloadAuth)

	usecase.NewLogin().Authenticate(bodyPayloadAuth.username, bodyPayloadAuth.password)

	if usecase.NewLogin().Authenticate(bodyPayloadAuth.username, bodyPayloadAuth.password) {
		g.JSON(http.StatusOK, gin.H{
			"message": "anda berhasil login", "data": bodyPayloadAuth,
		})
	} else {
		g.JSON(http.StatusUnauthorized, gin.H{
			"message": "anda gagal login", "data": err,
		})
	}
}
