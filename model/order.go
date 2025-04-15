package model

import (
	"time"
)

type Order struct {
	ID        string    `json:"id"`
	Customer  Customer  `json:"customer"`
	Items     []Item    `json:"items"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
