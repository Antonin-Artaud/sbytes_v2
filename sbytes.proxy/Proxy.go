package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type credential struct {
	Url        string `json:"url" binding:"required"`
	Credential struct {
		Id       string `json:"id" binding:"required"`
		Password string `json:"password" binding:"required"`
	} `json:"credential" binding:"required"`
}
type iProxy interface {
	handleAuthRequest(ctx *gin.RouterGroup)
}

type proxy struct {
	iProxy
}

func NewProxy() *proxy {
	return &proxy{}
}

func (receiver *proxy) handleAuthRequest(context *gin.RouterGroup) {
	context.POST("/login", func(ctx *gin.Context) {
		var data credential
		err := ctx.BindJSON(&data)

		if err != nil {
			log.Printf("somethings went wrong with post data...")
		}

		log.Printf("url is '%s', id is '%s', and password is '%s'",
			data.Url, data.Credential.Id, data.Credential.Password)
	})
}
