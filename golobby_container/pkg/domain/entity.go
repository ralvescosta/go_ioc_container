package domain

import "time"

type Entity struct {
	Id        int
	Name      string
	Email     string
	CreatedAt time.Time
}
