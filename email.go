package epostix

import "time"

type EmailCreate struct {
	From         string            `json:"from"`
	To           []string          `json:"to"`
	Subject      string            `json:"subject"`
	CC           *[]string         `json:"cc"`
	BCC          *[]string         `json:"bcc"`
	Text         *string           `json:"text"`
	HTML         *string           `json:"html"`
	Headers      map[string]string `json:"headers"`
	TemplateID   *string           `json:"templateID"`
	TemplateData map[string]string `json:"templateData"`
}

type Email struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
