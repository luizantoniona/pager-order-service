package model

import (
	"time"
)

type Order struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Customer  Customer  `json:"customer"`
	Address   Address   `json:"address"`
	Items     []Item    `json:"items"`
}
