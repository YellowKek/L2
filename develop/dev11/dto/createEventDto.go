package dto

type CreateEventDto struct {
	UserId int64  `json:"user_id"`
	Name   string `json:"name"`
	Date   string `json:"date"`
}
