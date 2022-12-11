package domain

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
	s.addBitrhdayPeople(people, birthday)
	return s.BirthdayPeople
}

func (s *Service) addBitrhdayPeople(people []models.Person, birthday models.DoB) {
	for _, person := range people {
		s.regularCase(person, birthday)
		s.leapYearCase(person, birthday)
	}
}

func (s *Service) regularCase(person models.Person, birthday models.DoB) {
	if person.Dob.Month == birthday.Month && person.Dob.Day == birthday.Day && person.Dob.Year < birthday.Year {
		s.BirthdayPeople = append(s.BirthdayPeople, person)
	}
}

func (s *Service) leapYearCase(person models.Person, birthday models.DoB) {
	if s.isLeapYear(person.Dob.Year) && person.Dob.Month == time.February && person.Dob.Day == 29 && person.Dob.Year < birthday.Year {
		if birthday.Month == time.February && birthday.Day == 28 {
			s.BirthdayPeople = append(s.BirthdayPeople, person)
		}
	}
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

func (s *Service) isLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
