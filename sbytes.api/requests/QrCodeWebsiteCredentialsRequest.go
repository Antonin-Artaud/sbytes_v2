package requests

import "time"

type (
	QrCodeWebsiteCredentialsRequest struct {
		Uri                string `json:"uri" binding:"required"`
		Signature          string `json:"signature"`
		TicketGuid         string `json:"ticketGuid"`
		CredentialRequired struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"credentialRequired" binding:"required"`
		ExpirationDelay time.Duration `json:"expirationDelay"`
	}
)
