//go:build unit

package src_test

import (
	"github.com/mikejeuga/bday/specifications"
	"github.com/mikejeuga/bday/src"
	"testing"
)

func TestService(t *testing.T) {
	service := src.NewService()
	birthdaySpec := specifications.NewGreetBirthday(service)
	birthdaySpec.SendGreetings(t)
}
