package domain

import (
	"context"
	"github.com/mikejeuga/bday/models"
	"time"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s Service) GreetBirthday(ctx context.Context, people []models.Person, today time.Time) []models.Person {
	birthday := s.timeToDoB(today)
	return s.addBitrhdayPeople(people, birthday)
}

func (s *Service) addBitrhdayPeople(people []models.Person, birthday models.DoB) []models.Person {
	var birthdayPeople []models.Person
	for _, person := range people {
		regularCases := s.regularCase(person, birthday)
		for _, regularCase := range regularCases {
			birthdayPeople = append(birthdayPeople, regularCase)
		}
		leapYearCases := s.leapYearCase(person, birthday)
		for _, leapYearCase := range leapYearCases {
			birthdayPeople = append(birthdayPeople, leapYearCase)
		}
	}
	return birthdayPeople
}

func (s *Service) regularCase(person models.Person, birthday models.DoB) []models.Person {
	var birthdayPeople []models.Person
	if person.Dob.Month == birthday.Month && person.Dob.Day == birthday.Day && person.Dob.Year < birthday.Year {
		birthdayPeople = append(birthdayPeople, person)
	}
	return birthdayPeople
}

func (s *Service) leapYearCase(person models.Person, birthday models.DoB) []models.Person {
	var leapYearPeople []models.Person
	if s.isLeapYear(person.Dob.Year) && person.Dob.Month == time.February && person.Dob.Day == 29 && person.Dob.Year < birthday.Year {
		if birthday.Month == time.February && birthday.Day == 28 {
			leapYearPeople = append(leapYearPeople, person)
		}
	}
	return leapYearPeople
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
