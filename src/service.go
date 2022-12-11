package src

import (
	"context"
	"github.com/mikejeuga/bday/models"
	"time"
)

type Service struct {
	BirthdayPeople []models.Person
}

func (s Service) GreetBirthday(ctx context.Context, people []models.Person, today time.Time) []models.Person {
	birthday := s.timeToDoB(today)
	for _, person := range people {
		if person.Dob.Month == birthday.Month && person.Dob.Day == birthday.Day && person.Dob.Year < birthday.Year {
			s.BirthdayPeople = append(s.BirthdayPeople, person)
		}
	}
	return s.BirthdayPeople
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
