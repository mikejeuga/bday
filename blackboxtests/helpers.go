package blackboxtests

import "net/http"

type WebTestClient struct {
	baseURL string
	Client  *http.Client
}
