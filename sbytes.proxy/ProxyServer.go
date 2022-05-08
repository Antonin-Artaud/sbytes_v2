package main

import (
	"github.com/gin-gonic/gin"
)

const (
	port      = ":8080"
	authRoute = "/auth"
)

func main() {
	r := gin.Default()

	authRequest := r.Group(authRoute)

	proxy := NewProxy()
	proxy.handleAuthRequest(authRequest)

	err := r.Run(port)

	if err != nil {
		panic("Somethings went wrong when proxy started...")
	}
}
