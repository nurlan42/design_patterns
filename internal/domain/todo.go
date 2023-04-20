package domain

import "time"

type TODO struct {
	ID          string    `json:"id,omitempty"`
	PublicID    string    `json:"publicID,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	DueDate     time.Time `json:"dueDate"`
	Completed   bool      `json:"completed,omitempty"`
}
