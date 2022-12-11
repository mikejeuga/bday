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

func (b *GreetBirthday) SendBirthdayGreetings(t *testing.T, today time.Time, people []models.Person) {
	s := testcase.NewSpec(t)
	ctx := testcase.Let(s, func(t *testcase.T) context.Context {
		return context.Background()
	})
	testPeople := testcase.Let(s, func(t *testcase.T) []models.Person {
		return people
	})
	s.Describe(".GreetBirthday", func(s *testcase.Spec) {

		act := func(t *testcase.T) []models.Person {
			return b.GreetBirthday(ctx.Get(t), testPeople.Get(t), today)
		}

		s.When("people in the list have their birthday today", func(s *testcase.Spec) {
			s.Then("they are greeted with an Happy Birthday message", func(t *testcase.T) {
				person := act(t)
				t.Must.Equal(testPeople.Get(t)[0], person[0])
			})
		})

		s.When("people in the list were born today", func(s *testcase.Spec) {
			BolBol := models.NewPerson("Bol", "Bol", models.DoB{
				Day:   4,
				Month: time.December,
				Year:  time.Now().Year(),
			})
			s.Before(func(t *testcase.T) {
				testPeople.Set(t, []models.Person{people[0], BolBol})
			})

			s.Then("they do not ", func(t *testcase.T) {
				t.Must.Equal(1, len(act(t)))
				t.Must.Equal(testPeople.Get(t)[0], act(t)[0])
			})
		})
	})

}
