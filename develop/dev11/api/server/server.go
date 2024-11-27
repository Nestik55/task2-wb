package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Nestik55/develop/dev11/api/service"
	"github.com/Nestik55/develop/dev11/api/service/repo"
)

const layout = "02-01-2006 15:04:05"

var business = service.NewService()

type Response struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

func getDataCreateDelete(r *http.Request) (repo.Event, error) {
	err := r.ParseForm()
	if err != nil {
		return repo.Event{}, err
	}

	event := repo.Event{}

	user_id := r.FormValue("user_id")
	userID, err := strconv.Atoi(user_id)
	if err != nil {
		return repo.Event{}, err
	}
	event.UserID = userID

	timeStr := r.FormValue("time")
	time, err := time.Parse(layout, timeStr)
	if err != nil {
		return repo.Event{}, err
	}
	event.Time = time
	event.Description = r.FormValue("description")

	return event, nil
}

func getDataUpdate(r *http.Request) ([2]repo.Event, error) {
	err := r.ParseForm()
	if err != nil {
		return [2]repo.Event{}, err
	}

	events := [2]repo.Event{}

	user_id := r.FormValue("user_id")
	userID, err := strconv.Atoi(user_id)
	if err != nil {
		return [2]repo.Event{}, err
	}
	events[0].UserID = userID
	events[1].UserID = userID

	timeStr := r.FormValue("new_time")
	newTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return [2]repo.Event{}, err
	}

	events[0].Time = newTime

	timeStr = r.FormValue("old_time")
	oldTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return [2]repo.Event{}, err
	}

	events[1].Time = oldTime

	events[0].Description = r.FormValue("new_description")
	events[1].Description = r.FormValue("old_description")

	return events, nil
}

func newResponse(w http.ResponseWriter, code int, result, text string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := Response{Result: result, Message: text}

	byteResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("newResponse: ", string(byteResponse), result, text)

	_, _ = w.Write(byteResponse)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newResponse(w, 400, "error", "Incorrect method")
		return
	}

	event, err := getDataCreateDelete(r)
	if err != nil {
		log.Print(err)
		newResponse(w, 400, "error", "Incorrect body")
		return
	}

	err = business.Create(&event)
	if err != nil {
		newResponse(w, 503, "error", "Error on service")
		return
	}

	newResponse(w, 200, "result", "Event created")
}

func updateEvent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		newResponse(w, 400, "error", "Incorrect method")
		return
	}

	events, err := getDataUpdate(r)
	if err != nil {
		newResponse(w, 400, "error", "Incorrect body")
		return
	}

	err = business.Update(&events[0], &events[1])
	if err != nil {
		newResponse(w, 503, "error", err.Error())
		return
	}

	newResponse(w, 200, "result", "Event updated")
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		newResponse(w, 400, "error", "Incorrect method")
		return
	}

	event, err := getDataCreateDelete(r)
	if err != nil {
		newResponse(w, 400, "error", "Incorrect body")
		return
	}

	err = business.Delete(&event)
	if err != nil {
		newResponse(w, 503, "error", "Error on service")
		return
	}

	newResponse(w, 200, "result", "Event deleted")
}

func getEventDay(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		newResponse(w, 400, "error", "Incorrect method")
		return
	}

	userIDstr := r.URL.Query().Get("user_id")
	if userIDstr == "" {
		newResponse(w, 400, "error", "Incorrect query string (have not id)")
	}

	userID, err := strconv.Atoi(userIDstr)

	if err != nil {
		newResponse(w, 400, "error", "Incorrect query string (have not id)")
		return
	}

	err, events := business.Get(userID, time.Now(), "d")
	if err != nil {
		newResponse(w, 503, "error", "Error on service")
	}

	res, err := json.Marshal(events)
	if err != nil {
		newResponse(w, 503, "error", "Error on service")
	}
	newResponse(w, 200, "result", string(res))
}

func getEventWeek(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		newResponse(w, 400, "error", "Incorrect method")
		return
	}

	userIDstr := r.URL.Query().Get("user_id")
	if userIDstr == "" {
		newResponse(w, 400, "error", "Incorrect query string (have not id)")
	}

	userID, err := strconv.Atoi(userIDstr)

	if err != nil {
		newResponse(w, 400, "error", "Incorrect query string (have not id)")
		return
	}

	err, events := business.Get(userID, time.Now(), "w")
	if err != nil {
		newResponse(w, 503, "error", "Error on service")
	}

	res, err := json.Marshal(events)
	if err != nil {
		newResponse(w, 503, "error", "Error on service")
	}
	newResponse(w, 200, "result", string(res))
}

func getEventMonth(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		newResponse(w, 400, "error", "Incorrect method")
		return
	}

	userIDstr := r.URL.Query().Get("user_id")
	if userIDstr == "" {
		newResponse(w, 400, "error", "Incorrect query string (have not id)")
	}

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		newResponse(w, 400, "error", "Incorrect query string (have not id)")
		return
	}

	err, events := business.Get(userID, time.Now(), "m")
	if err != nil {
		newResponse(w, 503, "error", "Error on service")
	}

	res, err := json.Marshal(events)
	if err != nil {
		newResponse(w, 503, "error", "Error on service")
	}
	newResponse(w, 200, "result", string(res))
}

func Run(host, port string) {

	mux := http.NewServeMux()

	mux.Handle("/create_event", loggingMiddleware(http.HandlerFunc(createEvent)))
	mux.Handle("/update_event", loggingMiddleware(http.HandlerFunc(updateEvent)))
	mux.Handle("/delete_event", loggingMiddleware(http.HandlerFunc(deleteEvent)))
	mux.Handle("/events_for_day", loggingMiddleware(http.HandlerFunc(getEventDay)))
	mux.Handle("/events_for_week", loggingMiddleware(http.HandlerFunc(getEventWeek)))
	mux.Handle("/events_for_month", loggingMiddleware(http.HandlerFunc(getEventMonth)))

	fmt.Println("server is listening..")
	err := http.ListenAndServe(host+":"+port, mux)
	log.Fatal(err)
}
