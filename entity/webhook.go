package entity

type WebhookPayload struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
	Key    string `json:"-"`
	URL    string `json:"-"`
}
