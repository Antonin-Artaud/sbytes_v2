package controllers

import (
	"github.com/gin-gonic/gin"
	"sbytes.api/requests"
)

type (
	iQrCodeHandler interface {
		NewQrCode(ctx *gin.Context)
	}

	qrCode struct {
		iQrCodeHandler
	}
)

func NewQrCode() *qrCode {
	return &qrCode{}
}

func (c *qrCode) CreateQrCode(ctx *gin.Context) {
	err := ctx.ShouldBindJSON(requests.QrCodeWebsiteCredentialsRequest{})
	if err != nil {
		ctx.JSON(400, gin.H{"message": "somethings went wrong..."})
	}
}
