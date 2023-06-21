package helpers

import (
	"io"
	"log"
	"net/http"
)

type ResponseData struct {
	Status int
	Url string
	Data string
}

type GetParams struct {
	Url string
	Client *http.Client
	Query map[string]string
}

type Fetch struct {}

// This executes a GET method
func (f *Fetch) Get(params GetParams) (data ResponseData, err error) {
	url, client, query := params.Url, params.Client, params.Query

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	for key, value := range query {
		query := req.URL.Query()
		query.Add(key, value)
		req.URL.RawQuery = query.Encode()
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil {
			log.Fatal(err)
		}
		
		bodyString := string(bodyBytes)
		
	return ResponseData{Status: res.StatusCode, Url: req.URL.String(), Data: bodyString}, nil
}