package main

import (
	"github.com/gin-gonic/gin"
	"sbytes.api/controllers"
	"sbytes.api/services"
)

const (
	port = ":8080"
)

func main() {

	server := gin.Default()
	services.GetInstance().InitiateDbConnection()

	qrCodeController := server.Group("/qrCodes")
	{
		/// TODO : refactoring this controller
		qrCode := controllers.NewQrCode()

		qrCodeController.GET("/", qrCode.CreateQrCode)
	}

	ticketsController := server.Group("/tickets")
	{
		ticketsHandler := controllers.NewTicketController()

		ticketsController.POST("/", ticketsHandler.Create)
		ticketsController.GET("/:uuid", ticketsHandler.ReadTicket)
		ticketsController.PUT("/:uuid", ticketsHandler.UpdateTicket)
	}

	err := server.Run(port)

	tryToHandleError(err)
}

func tryToHandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
