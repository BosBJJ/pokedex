package pokeapi

import (
	"net/http"
	"time"

	"github.com/BosBJJ/pokedex/internal/pokecache"
)

type PokeClient struct {
	Cache 	pokecache.Cache
	Client 	http.Client
}

func NewPokeClient(timeout, cacheInterval time.Duration) PokeClient {
	return PokeClient{
		Cache: pokecache.NewCache(cacheInterval),
		Client: http.Client{Timeout: timeout},
	}
}