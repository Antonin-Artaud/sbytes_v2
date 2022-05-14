package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

const (
	expirationDelay = 60 * time.Second
)

type (
	TicketController struct {
	}

	Ticket struct {
		Guid       uuid.UUID `json:"guid,omitempty"`
		Expiration int64     `json:"expiration,omitempty"`
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
		Guid:       newUUID,
		Expiration: time.Now().UTC().Add(expirationDelay).Unix(),
	}, nil
}

func (c *TicketController) Create(ctx *gin.Context) {

	ticket, err := c.NewTicket()

	if err != nil {
		ctx.JSON(422, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(201, gin.H{
		"ticket": &ticket,
	})

}

func (c *TicketController) ReadTicket(ctx *gin.Context) {
	ticket := &Ticket{}
	err := ctx.BindJSON(&ticket)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": errors.New("something went wrong when the ticket tried to be read"),
		})
	}

	if ticket.Expiration > time.Now().UTC().Unix() {
		ctx.JSON(410, gin.H{
			"error": errors.New("your ticket is expired, please create a new one"),
		})
	}

	ctx.JSON(200, gin.H{
		"ticket": &ticket,
	})

}

func (c *TicketController) UpdateTicket(ctx *gin.Context) {

}
