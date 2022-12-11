package blackboxtests

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/mikejeuga/bday/models"
	"io"
	"net/http"
	"time"
)

type WebTestClient struct {
	baseURL string
	Client  *http.Client
}

func NewWebTestClient() *WebTestClient {
	c := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   5 * time.Millisecond,
	}
	return &WebTestClient{Client: c}
}

func (c WebTestClient) GreetBirthday(ctx context.Context, people []models.Person, today time.Time) []models.Person {
	url := c.baseURL + "/birthdays"
	bytesPeople, err := json.Marshal(people)
	if err != nil {
		return []models.Person{}
	}
	jsonData := bytes.NewBuffer(bytesPeople)

	req, err := http.NewRequest(http.MethodGet, url, jsonData)
	if err != nil {
		return []models.Person{}
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return []models.Person{}
	}

	var birthdayPeople []models.Person

	returnedData, err := io.ReadAll(res.Body)
	if err != nil {
		return []models.Person{}
	}

	err = json.Unmarshal(returnedData, birthdayPeople)
	if err != nil {
		return []models.Person{}
	}

	return birthdayPeople

}
