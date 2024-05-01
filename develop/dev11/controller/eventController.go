package controller

import (
	"L2/develop/dev11/dto"
	"L2/develop/dev11/entity"
	"L2/develop/dev11/mapper"
	"L2/develop/dev11/repository"
	"encoding/json"
	"io"
	"net/http"
)

type EventController struct {
	EventRepo   *repository.EventRepo
	EventMapper *mapper.EventMapper
}

func NewEventController() *EventController {
	return &EventController{
		EventRepo:   repository.NewEventRepo(),
		EventMapper: mapper.NewEventMapper(),
	}
}

func (c *EventController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := c.getRequestBody(r)
		if err != nil {
			http.Error(w, "error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		eventToCreate := dto.CreateEventDto{}

		if err = json.Unmarshal(body, &eventToCreate); err != nil {
			http.Error(w, "error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		newEvent, err := c.EventMapper.CreateToEntity(eventToCreate)
		if err != nil {
			http.Error(w, "invalid data", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		newEvent = c.EventRepo.CreateEvent(newEvent)
		resp, err := json.Marshal(newEvent)
		if err != nil {
			http.Error(w, "error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(resp)

	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func (c *EventController) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := c.getRequestBody(r)
		if err != nil {
			http.Error(w, "error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		eventToUpdate := dto.UpdateEventDto{}
		if err := json.Unmarshal(body, &eventToUpdate); err != nil {
			http.Error(w, "error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		event, err := c.EventMapper.UpdateToEntity(eventToUpdate)
		if err != nil {
			http.Error(w, "invalid data", http.StatusBadRequest)
			return
		}

		updatedEvent, err := c.EventRepo.UpdateEvent(event)
		if err != nil {
			http.Error(w, "id не найден", http.StatusBadRequest)
			return
		}
		resp, err := json.Marshal(updatedEvent)
		if err != nil {
			http.Error(w, "error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resp)

	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func (c *EventController) getRequestBody(r *http.Request) ([]byte, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return body, nil
}

func (c *EventController) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		events := c.EventRepo.GetAll()
		var jsonEvents []byte
		for _, event := range events {
			temp, err := json.Marshal(event)
			if err != nil {
				http.Error(w, "error: "+err.Error(), http.StatusInternalServerError)
				return
			}
			for _, b := range temp {
				jsonEvents = append(jsonEvents, b)
			}
			jsonEvents = append(jsonEvents, "\n"[0])
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonEvents)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

func (c *EventController) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := c.getRequestBody(r)
		if err != nil {
			http.Error(w, "error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Id int64 `json:"id"`
		}{}

		if err = json.Unmarshal(body, &data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = c.EventRepo.DeleteById(data.Id); err != nil {
			http.Error(w, "id не найден", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

func (c *EventController) EventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		res := struct {
			result []entity.Event `json:"result"`
		}{}
		res.result = c.EventRepo.GetForDay()
		response, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(response)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

func (c *EventController) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		res := struct {
			result []entity.Event `json:"result"`
		}{}
		res.result = c.EventRepo.GetForWeek()
		response, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(response)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

func (c *EventController) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		res := struct {
			result []entity.Event `json:"result"`
		}{}
		res.result = c.EventRepo.GetForMonth()
		response, err := json.Marshal(res)
		if err != nil {
			http.Error(w, "error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(response)
	} else {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}
