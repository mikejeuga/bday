//go:build unit

package web_test

import (
	"bytes"
	"encoding/json"
	"github.com/adamluzsi/testcase"
	"github.com/adamluzsi/testcase/assert"
	"github.com/mikejeuga/bday/models"
	"github.com/mikejeuga/bday/src/adapters/web"
	"github.com/mikejeuga/bday/src/domain"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	s := testcase.NewSpec(t)

	s.Describe(".GET '/'", func(s *testcase.Spec) {
		var (
			resp    = httptest.NewRecorder()
			req     = httptest.NewRequest(http.MethodGet, "/", nil)
			service = domain.NewService()
			subject = testcase.Let(s, func(t *testcase.T) *http.Server {
				out := &bytes.Buffer{}
				return web.NewServer(service, out)
			})
		)

		act := func(t *testcase.T) {
			subject.Get(t).Handler.ServeHTTP(resp, req)
		}

		s.When("we call the server on '/',", func(s *testcase.Spec) {
			s.Then("the server is alive and running.", func(t *testcase.T) {
				act(t)
				t.Must.Equal(http.StatusOK, resp.Code)
			})
		})
	})

	s.Describe(".Get '/people?birthday=true'", func(s *testcase.Spec) {
		var (
			service = domain.NewService()
			out     = &bytes.Buffer{}
			subject = testcase.Let(s, func(t *testcase.T) *http.Server {
				return web.NewServer(service, out)
			})
		)
		var testPeople []models.Person
		johnDoe := models.NewPerson("John", "Doe", models.DoB{
			Year:  1982,
			Month: time.December,
			Day:   time.Now().Day(),
		})
		bruceWayne := models.NewPerson("Bruce", "Wayne", models.DoB{
			Year:  1982,
			Month: time.December,
			Day:   6,
		})
		stephCurry := models.NewPerson("Steph", "Curry", models.DoB{
			Year:  1965,
			Month: time.December,
			Day:   7,
		})
		ladyGaga := models.NewPerson("Lady", "Gaga", models.DoB{
			Year:  1986,
			Month: time.December,
			Day:   8,
		})

		testPeople = append(testPeople, johnDoe, stephCurry, bruceWayne, ladyGaga)

		marshal, err := json.Marshal(testPeople)
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/people?birthday=true", bytes.NewReader(marshal))
		req.Header.Set("Content-type", "application/json")

		act := func(t *testcase.T) {
			subject.Get(t).Handler.ServeHTTP(resp, req)
		}

		s.When("we call the server on '/people?birthday=true',", func(s *testcase.Spec) {
			s.Then("the server greeting the birthday.", func(t *testcase.T) {
				act(t)
				t.Must.Equal(http.StatusOK, resp.Code)
				t.Must.Equal("Happy Birthday John", out.String())
			})
		})
	})
}
