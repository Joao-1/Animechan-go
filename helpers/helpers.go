package helpers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type ResponseData struct {
	Status int
	Data string
}

type Fetch struct {}

// This executes a GET method
func (f *Fetch) Get(url string, client *http.Client) (data ResponseData, err error) {
	req, err := http.NewRequest("GET", url, nil)
	
	if err != nil {
		log.Fatal(err)
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
		
	return ResponseData{Status: res.StatusCode, Data: bodyString}, nil
}

type MockClient struct {
	mock.Mock
}

func (m *MockClient) RoundTrip(req *http.Request) (*http.Response, error) {
	args := m.Called(req)

	return args.Get(0).(*http.Response), args.Error(1)
}

func MockHttpClient() (*http.Client, *MockClient) {
	client := &http.Client{Transport: new(MockClient)}

	return client, client.Transport.(*MockClient)
}

func Server() {
	http.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		err := json.NewEncoder(res).Encode("pong")
		
		if err != nil {
			log.Fatal(err)
		}
	})

	go func() {
		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
}
