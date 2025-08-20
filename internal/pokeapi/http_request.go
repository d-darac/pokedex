package pokeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func sendHttpRequest(pokeClient *PokeClient, method, url string, body interface{}, headers *http.Header) (*http.Response, error) {
	buffer := bytes.Buffer{}

	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error while marshaling the body: %v", err)
		}

		buffer = *bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, &buffer)
	if err != nil {
		return nil, fmt.Errorf("error while making the new request: %v", err)
	}

	if headers != nil {
		for key, val := range *headers {
			for i := range val {
				req.Header.Set(key, val[i])
			}
		}
	}

	res, err := pokeClient.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing the request: %v", err)
	}

	return res, nil
}
