package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mikejeuga/bday/models"
	"github.com/mikejeuga/bday/specifications"
	"io"
	"net/http"
	"time"
)

type Server struct {
	out     io.Writer
	greeter specifications.BirthdayGreeter
}

func NewServer(greeter specifications.BirthdayGreeter, out io.Writer) *http.Server {
	router := mux.NewRouter()
	server := Server{out: out, greeter: greeter}

	router.HandleFunc("/", server.Home).Methods(http.MethodGet)
	router.HandleFunc("/birthdays", server.GreetBirthdays).Methods(http.MethodPost)

	return &http.Server{
		Addr:    ":8082",
		Handler: router,
	}
}

func (s Server) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is up and running")
}

func (s Server) GreetBirthdays(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var receivedPeople []models.Person
	err := json.NewDecoder(r.Body).Decode(&receivedPeople)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	birthdays := s.greeter.GreetBirthday(ctx, receivedPeople, time.Now())
	for _, birthday := range birthdays {
		fmt.Fprintf(s.out, fmt.Sprintf("Happy Birthday %s", birthday.FName))
	}
}
