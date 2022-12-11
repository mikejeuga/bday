//go:build acceptance

package blackboxtests

import (
	"github.com/mikejeuga/bday/specifications"
	"testing"
)

func TestAPI(t *testing.T) {
	t.Skip("let's wait for a server!")
	testClient := NewWebTestClient()
	greetingBirthdaysSpec := specifications.NewGreetBirthday(testClient)
	greetingBirthdaysSpec.SendGreetings(t)
}
