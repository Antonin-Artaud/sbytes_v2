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

type Ticket struct {
	Guid       uuid.UUID `json:"guid,omitempty"`
	Expiration int64     `json:"expiration,omitempty"`
}

func NewTicket() (*Ticket, error) {
	newUUID, err := uuid.NewUUID()

	if err != nil {
		return nil, errors.New("somethings went wrong when the ticket tried to be created")
	}

	return &Ticket{
		Guid:       newUUID,
		Expiration: time.Now().UTC().Add(expirationDelay).Unix(),
	}, nil
}

func (c *Ticket) Create(ctx *gin.Context) {

	ticket, err := NewTicket()

	if err != nil {
		ctx.JSON(422, gin.H{"error": err.Error()})
	}

	ctx.JSON(201, gin.H{
		"ticket": &ticket,
	})

}

func (c *Ticket) ReadTicket(ctx *gin.Context) {

}

func (c *Ticket) UpdateTicket(ctx *gin.Context) {

}
