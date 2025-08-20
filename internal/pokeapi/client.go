package pokeapi

import (
	"net/http"
	"time"
)

type PokeClient struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) PokeClient {
	return PokeClient{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
