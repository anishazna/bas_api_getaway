package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionInterface interface {
	TransferBank(*gin.Context)
}

type transactionImplement struct{}

func NewTransaction() transactionInterface {
	return &transactionImplement{}
}

type bodyPayloadTransaction struct{}

func (a *transactionImplement) TransferBank(g *gin.Context) {
	bodyPayloadTxn := bodyPayloadTransaction{}
	err := g.BindJSON(&bodyPayloadTxn)

	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Hello guys this API rest for later",
	})
}
