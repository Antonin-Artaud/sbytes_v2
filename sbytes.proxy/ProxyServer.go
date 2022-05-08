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

	handleError(err, "Somethings went wrong when the server has started...")
}

func handleError(err error, message string) {
	if err != nil {
		panic(message)
	}
}