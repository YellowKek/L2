package repository

import (
	"L2/develop/dev11/entity"
	"errors"
	"fmt"
	"time"
)

type EventRepo struct {
	eventId  int64
	eventMap map[int64]entity.Event
}

func NewEventRepo() *EventRepo {
	return &EventRepo{eventId: 1, eventMap: make(map[int64]entity.Event)}
}

func (r *EventRepo) GetById(id int64) (entity.Event, error) {
	v, ok := r.eventMap[id]
	if ok {
		return v, nil
	}
	return entity.Event{}, errors.New("id не найден")
}

func (r *EventRepo) CreateEvent(event entity.Event) entity.Event {
	event.Id = r.eventId
	r.eventMap[r.eventId] = event
	r.eventId++
	return r.eventMap[event.Id]
}

func (r *EventRepo) UpdateEvent(event entity.Event) (entity.Event, error) {
	if _, err := r.GetById(event.Id); err != nil {
		return entity.Event{}, err
	}

	event.UserId = r.eventMap[event.Id].UserId
	r.eventMap[event.Id] = event
	return event, nil
}

func (r *EventRepo) DeleteById(id int64) error {
	_, err := r.GetById(id)
	if err == nil {
		delete(r.eventMap, id)
	}
	return err
}

func (r *EventRepo) GetAll() []entity.Event {
	events := make([]entity.Event, 0, len(r.eventMap))
	fmt.Println(r.eventMap, len(r.eventMap))
	for _, event := range r.eventMap {
		events = append(events, event)
	}
	return events
}

func (r *EventRepo) GetForDay() []entity.Event {
	events := make([]entity.Event, 0)
	for _, event := range r.eventMap {
		if event.Date.Sub(time.Now())/time.Hour <= 24 {
			events = append(events, event)
		}
	}
	return events
}

func (r *EventRepo) GetForWeek() []entity.Event {
	events := make([]entity.Event, 0)
	for _, event := range r.eventMap {
		if event.Date.Sub(time.Now())/(time.Hour*24) <= 7 {
			events = append(events, event)
		}
	}
	return events
}

func (r *EventRepo) GetForMonth() []entity.Event {
	events := make([]entity.Event, 0)
	for _, event := range r.eventMap {
		if event.Date.Sub(time.Now())/(time.Hour*24*7) <= 30 {
			events = append(events, event)
		}
	}
	return events
}
