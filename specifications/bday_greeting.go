package specifications

import (
	"context"
	"github.com/adamluzsi/testcase"
	"github.com/mikejeuga/bday/models"
	"testing"
	"time"
)

type BirthdayGreeter interface {
	GreetBirthday(ctx context.Context, person []models.Person, today time.Time) []models.Person
}

type GreetBirthday struct {
	BirthdayGreeter
}

func NewGreetBirthday(birthdayGreeter BirthdayGreeter) *GreetBirthday {
	return &GreetBirthday{BirthdayGreeter: birthdayGreeter}
}

func (g *GreetBirthday) SendGreetings(t *testing.T) {
	s := testcase.NewSpec(t)
	ctx := testcase.Let(s, func(t *testcase.T) context.Context {
		return context.Background()
	})

	today := testcase.Let[time.Time](s, nil)

	people := testcase.Let(s, func(t *testcase.T) []models.Person {
		jeffGreen := models.NewPerson("Jeff", "Green", models.DoB{
			Day:   4,
			Month: time.December,
			Year:  1999,
		})

		return []models.Person{jeffGreen}
	})
	act := func(t *testcase.T) []models.Person {
		return g.GreetBirthday(ctx.Get(t), people.Get(t), today.Get(t))
	}
	s.Describe(".GreetBirthday", func(s *testcase.Spec) {
		s.When("people in the list have their birthday today,", func(s *testcase.Spec) {

			s.Before(func(t *testcase.T) {
				today.Set(t, time.Date(time.Now().Year(), time.December, 4, 0, 0, 0, 0, time.UTC))
			})

			s.Then("they are selected to receive a Happy Birthday message.", func(t *testcase.T) {
				thePeople := act(t)
				t.Must.Equal(people.Get(t)[0], thePeople[0])
			})
		})

		s.When("people in the list were born today,", func(s *testcase.Spec) {
			BolBol := models.NewPerson("Bol", "Bol", models.DoB{
				Day:   4,
				Month: time.December,
				Year:  time.Now().Year(),
			})
			s.Before(func(t *testcase.T) {
				people.Set(t, []models.Person{people.Get(t)[0], BolBol})
				today.Set(t, time.Date(time.Now().Year(), time.December, 4, 0, 0, 0, 0, time.UTC))
			})

			s.Then("they are not shorlisted to receive a Happy Birthday message.", func(t *testcase.T) {
				t.Must.Equal(1, len(act(t)))
				t.Must.Equal(people.Get(t)[0], act(t)[0])
			})
		})

		s.When("a person selected in the list was born Feb 29 in a leap year,", func(s *testcase.Spec) {
			LipYear := models.NewPerson("Lip", "Hier", models.DoB{
				Day:   29,
				Month: time.February,
				Year:  1984,
			})
			s.Before(func(t *testcase.T) {
				people.Set(t, []models.Person{people.Get(t)[0], LipYear})
				today.Set(t, time.Date(time.Now().Year(), time.February, 28, 0, 0, 0, 0, time.UTC))
			})
			s.Then("receive a Happy Birthday message on Feb 28 of non leap years", func(t *testcase.T) {
				bdayPeople := act(t)
				t.Must.Equal(1, len(bdayPeople))
				t.Must.Equal(LipYear.FName, bdayPeople[0].FName)
			})
		})
	})
}
