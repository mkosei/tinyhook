package model

import "time"

type Event struct {
	ID        string
	Provider  string
	Headers   map[string]string
	Body      []byte
	CreatedAt time.Time
}
