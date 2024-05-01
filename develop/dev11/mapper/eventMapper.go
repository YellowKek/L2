package mapper

import (
	"L2/develop/dev11/dto"
	"L2/develop/dev11/entity"
	"time"
)

type EventMapper struct {
}

func NewEventMapper() *EventMapper {
	return &EventMapper{}
}

func (m *EventMapper) CreateToEntity(dto dto.CreateEventDto) (entity.Event, error) {
	date, err := time.Parse(time.DateOnly, dto.Date)
	if err != nil {
		return entity.Event{}, err
	}

	return entity.Event{UserId: dto.UserId, Name: dto.Name, Date: date}, nil
}

func (m *EventMapper) UpdateToEntity(dto dto.UpdateEventDto) (entity.Event, error) {
	date, err := time.Parse(time.DateOnly, dto.Date)
	if err != nil {
		return entity.Event{}, err
	}
	return entity.Event{Id: dto.Id, Name: dto.Name, Date: date}, nil
}
