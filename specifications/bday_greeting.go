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
	s.Describe(".GreetBirthday", func(s *testcase.Spec) {

		act := func(t *testcase.T) []models.Person {
			return b.GreetBirthday(ctx.Get(t), people, today)
		}

		s.When("people in the list have their birthday today", func(s *testcase.Spec) {
			s.Then("they are greeted with an Happy Birthday message", func(t *testcase.T) {
				person := act(t)
				t.Must.Equal(people[0], person[0])
			})
		})
	})

}
