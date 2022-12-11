//go:build unit

package src_test

import (
	"github.com/mikejeuga/bday/models"
	"github.com/mikejeuga/bday/specifications"
	"github.com/mikejeuga/bday/src"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	service := src.NewService()

	greetBirthdaySpec := specifications.NewGreetBirthday(service)
	today := time.Date(time.Now().Year(), time.December, 4, 0, 0, 0, 0, time.UTC)
	jeffGreen := models.NewPerson("Jeff", "Green", models.DoB{
		Day:   4,
		Month: time.December,
		Year:  1999,
	})
	people := []models.Person{jeffGreen}
	greetBirthdaySpec.SendBirthdayGreetings(t, today, people)
}
