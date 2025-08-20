package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (pokeClient *PokeClient) ListLocationAreas(pageUrl *string) (PokeAPIListResponse, error) {
	url := baserUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	res, err := sendHttpRequest(pokeClient, http.MethodGet, url, nil, nil)
	if err != nil {
		return PokeAPIListResponse{}, err
	}
	defer res.Body.Close()

	var listRes PokeAPIListResponse

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&listRes); err != nil {
		return PokeAPIListResponse{}, fmt.Errorf("error while decoding json data: %v", err)
	}

	return listRes, nil
}
