package src

import (
	"context"
	"fmt"
	"github.com/mikejeuga/bday/models"
	"time"
)

type Service struct{}

func (s Service) GreetBirthday(ctx context.Context, people []models.Person, today time.Time) ([]models.Person, error) {
	birthday := s.timeToDoB(today)
	if people[0].Dob.Month == birthday.Month && people[0].Dob.Day == birthday.Day {
		return []models.Person{people[0]}, nil
	}
	return []models.Person{}, fmt.Errorf("Not a single person has there birthday today.")
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
