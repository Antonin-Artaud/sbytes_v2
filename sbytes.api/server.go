package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sbytes.api/controllers"
	"sbytes.api/services"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		throwCriticalError(err)
	}

	setServerMode(os.Getenv("APP_ENV"))

	server := gin.Default()

	if err := server.SetTrustedProxies([]string{os.Getenv("IP_TRUSTED")}); err != nil {
		throwCriticalError(err)
	}

	if err := services.GetService().InitiateDbConnection(); err != nil {
		throwCriticalError(err)
	}

	ticketsController := server.Group("/ticket")
	{
		ticketsHandler := controllers.NewTicketController()

		ticketsController.POST("/", ticketsHandler.Create)
		ticketsController.GET("/:uuid", ticketsHandler.ReadTicket)
		ticketsController.PUT("/:uuid", ticketsHandler.UpdateTicket)
	}

	if err := server.Run(os.Getenv("SERVER_PORT")); err != nil {
		throwCriticalError(err)
	}
}

func setServerMode(mode string) {
	switch mode {
	case "DEV":
		gin.SetMode(gin.DebugMode)
		break
	case "PROD":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}

func throwCriticalError(err error) {
	log.Println(err.Error())
	return
}
