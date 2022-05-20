package responses

import "time"

type (
	CreateTicketResponse struct {
		TicketGuid interface{}   `json:"TicketGuid"`
		Timeout    time.Duration `json:"Timeout"`
	}

	WebsiteSimpleCredentialsResponse struct {
		Identifier string `json:"Identifier"`
		Password   string `json:"Password"`
	}

	WebSiteCompleteCredentialsResponse struct {
		Identifier string `json:"Identifier"`
		Password   string `json:"Password"`
	}
)
