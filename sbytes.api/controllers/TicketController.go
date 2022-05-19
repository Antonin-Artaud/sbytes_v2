package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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
		Uuid       string `json:"guid,omitempty" bson:"guid,omitempty"`
		Expiration int64  `json:"expiration,omitempty" bson:"expiration,omitempty"`
	}
)

func NewTicketController() *TicketController {
	return &TicketController{}
}

func (c *TicketController) NewTicket() (*Ticket, error) {
	newUUID, err := uuid.NewUUID()

	if err != nil {
		return nil, errors.New("somethings went wrong when the ticket tried to be created")
	}

	return &Ticket{
		Uuid:       newUUID.String(),
		Expiration: time.Now().UTC().Add(expirationDelay).Unix(),
	}, nil
}

func (c *TicketController) Create(ctx *gin.Context) {

	ticket, err := c.NewTicket()

	doc := bson.D{{"ticket", ticket}}

	if err != nil {
		ctx.JSON(422, err.Error())
		return
	}

	err = services.GetInstance().MongoDb.InsertTicket(doc)

	if err != nil {
		ctx.JSON(422, err.Error())
		return
	}

	if err != nil {
		ctx.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"ticket": &ticket,
	})
}

func (c *TicketController) ReadTicket(ctx *gin.Context) {
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

func (c *TicketController) UpdateTicket(ctx *gin.Context) {

}
