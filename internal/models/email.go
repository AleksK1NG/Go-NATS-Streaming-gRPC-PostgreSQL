package models

import "time"

type Email struct {
	EmailID   uint64    `json:"emailID"`
	From      string    `json:"from"`
	To        []string  `json:"to"`
	CreatedAt time.Time `json:"createdAt"`
}
