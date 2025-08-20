package pokeapi

import (
	"net/http"
	"time"

	"github.com/d-darac/pokedex/internal/pokecache"
)

type PokeClient struct {
	cache      pokecache.PokeCache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) PokeClient {
	return PokeClient{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
