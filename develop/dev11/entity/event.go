package entity

import "time"

type Event struct {
	Id     int64     `json:"id"`
	UserId int64     `json:"user_id"`
	Name   string    `json:"name"`
	Date   time.Time `json:"date"`
}
