package controllers

import "github.com/gin-gonic/gin"

type Ticket struct{}

func NewTicket() *Ticket {
	return &Ticket{}
}

func (c *Ticket) Create(ctx *gin.Context) {

}

func (c *Ticket) ReadTicket(ctx *gin.Context) {
	
}

func (c *Ticket) UpdateTicket(ctx *gin.Context) {

}