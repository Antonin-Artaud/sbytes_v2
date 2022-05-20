package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		Uuid       string `json:"guid" bson:"guid"`
		Expiration int64  `json:"expiration" bson:"expiration"`
	}
)

func NewTicketController() *TicketController {
	return &TicketController{}
}

func (receiver *TicketController) NewTicket() (*Ticket, error) {
	newUUID, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	return &Ticket{
		Uuid:       newUUID.String(),
		Expiration: time.Now().UTC().Add(expirationDelay).Unix(),
	}, nil
}

func (receiver *TicketController) Create(ctx *gin.Context) {
	var req requests.WebsiteCredentialsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		receiver.sendHttpResponse(ctx, 500, gin.H{"error": err.Error()})
		return
	}

	ticket, err := receiver.NewTicket()

	if err != nil {
		receiver.sendHttpResponse(ctx, 500, gin.H{"error": err.Error()})
		return
	}

	var bsonDocument bson.D

	bsonDocument = bson.D{{"ticket", ticket}}

	if err := services.GetInstance().MongoDb.InsertTicket(bsonDocument); err != nil {
		receiver.sendHttpResponse(ctx, 500, gin.H{"error": err.Error()})
		return
	}

	response := &responses.CreateTicketResponse{
		TicketGuid: ticket.Uuid,
		Timeout:    60 * time.Second,
	}

	receiver.sendHttpResponse(ctx, 201, gin.H{"response": response})
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
