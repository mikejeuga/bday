package specifications

import (
	"context"
	"github.com/adamluzsi/testcase"
	"github.com/mikejeuga/bday/models"
	"testing"
	"time"
)

type BirthdayGreeter interface {
	GreetBirthday(ctx context.Context, person []models.Person, today time.Time) ([]models.Person, error)
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

		act := func(t *testcase.T) ([]models.Person, error) {
			return b.GreetBirthday(ctx.Get(t), people, today)
		}

		s.Then("each birthday person is greeted with an Happy Birthday message", func(t *testcase.T) {
			person, err := act(t)
			t.Must.NoError(err)
			t.Must.Equal(people[0], person[0])
		})
	})

}
