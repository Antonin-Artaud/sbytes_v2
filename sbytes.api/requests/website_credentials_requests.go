package requests

type (
	WebsiteCredentialsRequest struct {
		Uri              string `json:"Uri"`
		FormatCredential string `json:"FormatCredential"`
	}
)
