package main

import (
	"L2/develop/dev11/controller"
	"log"
	"net/http"
	"time"
)

func main() {
	eventController := controller.NewEventController()

	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", eventController.Create)
	mux.HandleFunc("/update_event", eventController.Update)
	mux.HandleFunc("/delete_event", eventController.Delete)
	mux.HandleFunc("/events_for_day", eventController.EventsForDay)
	mux.HandleFunc("/events_for_week", eventController.EventsForWeek)
	mux.HandleFunc("/events_for_month", eventController.EventsForMonth)

	http.HandleFunc("/get_all", eventController.GetAll)

	logging := loggingMiddleware(mux)
	err := http.ListenAndServe(":8080", logging)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}
