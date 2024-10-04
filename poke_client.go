package main

import (
	"net/http"
	"time"
)

type PokeClient struct {
	httpClient http.Client
}

func newClient(timeout time.Duration) PokeClient {
	return PokeClient{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
