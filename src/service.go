package src

import (
	"context"
	"github.com/mikejeuga/bday/models"
	"time"
)

type Service struct{}

func (s Service) GreetBirthday(ctx context.Context, people []models.Person, today time.Time) []models.Person {
	var birthdayPeople []models.Person
	birthday := s.timeToDoB(today)
	for _, person := range people {
		if person.Dob.Month == birthday.Month && person.Dob.Day == birthday.Day {
			birthdayPeople = append(birthdayPeople, person)
		}
	}
	return birthdayPeople
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) timeToDoB(date time.Time) models.DoB {
	return models.DoB{
		Day:   date.Day(),
		Month: date.Month(),
		Year:  date.Year(),
	}
}
