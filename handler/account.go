package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountInterface interface {
	GetAccount(*gin.Context)
	CreateAccount(*gin.Context)
	UpdateAccount(*gin.Context)
	DeleteAccount(*gin.Context)
	BalanceAccount(*gin.Context)
}

type accountImplement struct{}

type BodyPayloadBalance struct{}

func NewAccount() AccountInterface {
	return &accountImplement{}
}

func (a *accountImplement) GetAccount(g *gin.Context) {
	queryParam := g.Request.URL.Query()

	name := queryParam.Get("name")

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    name,
	})
}

type BodyPayloadAccount struct {
	AccountID string
	Name      string
	Address   string
}

func (a *accountImplement) CreateAccount(g *gin.Context) {
	bodyPayload := BodyPayloadAccount{}

	err := g.BindJSON(&bodyPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    bodyPayload,
	})
}
func (a *accountImplement) UpdateAccount(g *gin.Context) {
	queryParam := g.Request.URL.Query()

	name := queryParam.Get("name")

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    name,
	})
}

func (a *accountImplement) BalanceAccount(g *gin.Context) {
	bodyPayloadBalance := BodyPayloadBalance{}
	err := g.BindJSON(&bodyPayloadBalance)

	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Hello guys this API rest for later",
	})
}

func (a *accountImplement) DeleteAccount(g *gin.Context) {
	id := g.Param("id")

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    id,
	})
}
