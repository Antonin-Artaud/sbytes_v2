package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"sbytes.api/controllers"
	"sbytes.api/services"
)

const (
	port = ":8080"
)

func main() {
	loadEnv()

	server := gin.Default()
	err := services.GetService().InitiateDbConnection()

	if err != nil {
		panic(err.Error())
	}

	ticketsController := server.Group("/ticket")
	{
		ticketsHandler := controllers.NewTicketController()

		ticketsController.POST("/", ticketsHandler.Create)
		ticketsController.GET("/:uuid", ticketsHandler.ReadTicket)
		ticketsController.PUT("/:uuid", ticketsHandler.UpdateTicket)
	}

	err = server.Run(port)

	if err != nil {
		panic(err.Error())
	}
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err.Error())
	}
}
