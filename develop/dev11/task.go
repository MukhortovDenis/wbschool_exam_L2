package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

// Server для работы с методами
type Server struct {
	srv *http.Server
}

type NewHandler struct{}

type Error struct {
	Err error `json:"error"`
}

type Result struct {
	Result UserEvent `json:"result"`
}

type UserEvent struct {
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
}

var UserMap = make(map[int]*UserEvent)

var month = [12]string{"January", "February", "March", "April", "May",
	"June", "July", "August", "September", "October", "Novembe", "December"}

var weekday = [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (s *Server) Run(handler http.Handler) error {
	s.srv = &http.Server{
		Handler: handler,
		Addr:    ":8080"}
	return s.srv.ListenAndServe()
}

func middleware(r *http.Request) {
	log.Println(r.URL.Path)
}

func (h NewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	middleware(r)
	switch r.URL.Path {
	case "/create_event":
		if r.Method == "POST" {
			createEvent(w, r)
		} else {
			w.WriteHeader(500)
		}
	case "/update_event":
		if r.Method == "POST" {
			updateEvent(w, r)
		} else {
			w.WriteHeader(500)
		}
	case "/delete_event":
		if r.Method == "POST" {
			deleteEvent(w, r)
		} else {
			w.WriteHeader(500)
		}
	case "/events_for_day":
		if r.Method == "GET" {
			eventForDay(w, r)
		} else {
			w.WriteHeader(500)
		}
	case "/events_for_week":
		if r.Method == "GET" {
			eventForWeek(w, r)
		} else {
			w.WriteHeader(500)
		}
	case "/events_for_month":
		if r.Method == "GET" {
			eventForMonth(w, r)
		} else {
			w.WriteHeader(500)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (u *UserEvent) GetResult(w http.ResponseWriter) {
	body := new(bytes.Buffer)
	res := Result{UserEvent{
		UserID: u.UserID,
		Date:   u.Date}}
	err := json.NewEncoder(body).Encode(res)
	if err != nil {
		GetError(err, w)
	}
	fmt.Fprint(w, body)
}

func GetError(err error, w http.ResponseWriter) {
	var newError Error
	newError.Err = err
	body := new(bytes.Buffer)
	err = json.NewEncoder(body).Encode(newError)
	if err != nil {
		log.Print(err)
	}
	fmt.Fprint(w, body)
	w.WriteHeader(503)
}

func (u *UserEvent) isValid(w http.ResponseWriter, r *http.Request) bool {
	var err error
	u.UserID, err = strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		GetError(err, w)
		w.WriteHeader(400)
		return false
	}
	u.Date = r.URL.Query().Get("date")
	if !strings.Contains(u.Date, "-") {
		err = errors.New("wrong date")
		GetError(err, w)
		w.WriteHeader(400)
		return false
	}
	return true
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var u UserEvent
	if u.isValid(w, r) {
		UserMap[u.UserID] = &u
		u.GetResult(w)
	}
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	var u UserEvent
	if u.isValid(w, r) {
		for _, user := range UserMap {
			if user.UserID == u.UserID {
				user.Date = u.Date
				user.GetResult(w)
				return
			}
		}
		err := errors.New("no record")
		GetError(err, w)
	}
}
func deleteEvent(w http.ResponseWriter, r *http.Request) {
	var u UserEvent
	if u.isValid(w, r) {
		for _, user := range UserMap {
			if user.UserID == u.UserID && user.Date == u.Date {
				delete(UserMap, u.UserID)
				fmt.Fprint(w, `{"result":"delete"}`)
				return
			}
		}
		err := errors.New("no record")
		GetError(err, w)
	}
}

func eventForDay(w http.ResponseWriter, r *http.Request) {
	for _, user := range UserMap {
		slice := strings.Split(user.Date, "-")
		a, err := strconv.Atoi(slice[0])
		if err != nil {
			GetError(err, w)
			return
		}
		b, err := strconv.Atoi(slice[1])
		if err != nil {
			GetError(err, w)
			return
		}
		c, err := strconv.Atoi(slice[2])
		if err != nil {
			GetError(err, w)
			return
		}
		if a == time.Now().Day() && month[b-1] == time.Now().Month().String() && c == time.Now().Year() {
			user.GetResult(w)
		}
	}
}

func eventForWeek(w http.ResponseWriter, r *http.Request) {
	d := time.Now().Weekday().String()
	for i, day := range weekday {
		if d == day {
			for _, user := range UserMap {
				slice := strings.Split(user.Date, "-")
				a, err := strconv.Atoi(slice[0])
				if err != nil {
					GetError(err, w)
					return
				}
				b, err := strconv.Atoi(slice[1])
				if err != nil {
					GetError(err, w)
					return
				}
				c, err := strconv.Atoi(slice[2])
				if err != nil {
					GetError(err, w)
					return
				}
				if a <= time.Now().Day() && a >= time.Now().Day()-i && month[b-1] == time.Now().Month().String() && c == time.Now().Year() {
					user.GetResult(w)
				}
			}
			return
		}
	}
}

func eventForMonth(w http.ResponseWriter, r *http.Request) {
	month := [12]string{"January", "February", "March", "April", "May",
		"June", "July", "August", "September", "October", "Novembe", "December"}
	for _, user := range UserMap {
		slice := strings.Split(user.Date, "-")
		b, err := strconv.Atoi(slice[1])
		if err != nil {
			GetError(err, w)
			return
		}
		c, err := strconv.Atoi(slice[2])
		if err != nil {
			GetError(err, w)
			return
		}
		if month[b-1] == time.Now().Month().String() && c == time.Now().Year() {
			user.GetResult(w)
		}
	}
}

func main() {
	handler := new(NewHandler)
	srv := new(Server)
	err := srv.Run(*handler)
	if err != nil {
		log.Fatal(err)
	}
}
