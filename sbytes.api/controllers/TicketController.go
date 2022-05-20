package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"sbytes.api/requests"
	"sbytes.api/responses"
	"sbytes.api/services"
	"time"
)

const (
	expirationDelay = 60 * time.Second
)

type (
	TicketController struct {
	}

	Ticket struct {
		Expiration int64 `json:"expiration" bson:"expiration"`
	}
)

func NewTicketController() *TicketController {
	return &TicketController{}
}

func (receiver *TicketController) NewTicket() *Ticket {
	return &Ticket{
		Expiration: time.Now().UTC().Add(expirationDelay).Unix(),
	}
}

func (receiver *TicketController) Create(ctx *gin.Context) {
	var req requests.WebsiteCredentialsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		receiver.sendHttpResponse(ctx, 500, gin.H{"error": err.Error()})
		return
	}

	ticket := receiver.NewTicket()

	bsonElement := bson.M{"expiration": ticket.Expiration}

	if insertedObjectId, err := services.GetInstance().MongoDb.InsertTicket(bsonElement); err != nil {
		receiver.sendHttpResponse(ctx, 500, gin.H{"error": err.Error()})
	} else {
		response := &responses.CreateTicketResponse{
			TicketGuid: insertedObjectId,
			Timeout:    60 * time.Second,
		}

		receiver.sendHttpResponse(ctx, 201, response)
	}
}

func (receiver *TicketController) ReadTicket(ctx *gin.Context) {
	guid := ctx.Param("uuid")
	ticket := services.GetInstance().MongoDb.FindTicket(guid)

	if ticket == nil {
		ctx.JSON(404, gin.H{
			"error": "Ticket not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"ticket": ticket[1],
	})
}

func (receiver *TicketController) UpdateTicket(ctx *gin.Context) {

}

func (receiver *TicketController) sendHttpResponse(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, data)
}
