package model

import (
	"time"
)

type Pager struct {
	ID          string    `json:"id"`
	PagerNumber string    `json:"pager_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Customer    *Customer `json:"customer,omitempty"`
	Items       []Item    `json:"items"`
}
