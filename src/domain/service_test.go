//go:build unit

package domain_test

import (
	"github.com/mikejeuga/bday/specifications"
	"github.com/mikejeuga/bday/src/domain"
	"testing"
)

func TestService(t *testing.T) {
	service := domain.NewService()
	birthdaySpec := specifications.NewGreetBirthday(service)
	birthdaySpec.SendGreetings(t)
}
