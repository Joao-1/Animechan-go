package helpers

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T){
	baseURL := "https://localhost:3000/"

	body := "Hello, World"
	client := &http.Client{}

	urlWithQuery := baseURL + "?foo=bar"

	defer gock.Off()
	gock.New(baseURL).Get("/").Reply(200).JSON(body)
	gock.New(urlWithQuery).Get("/").Reply(200).JSON(body)

	t.Run("Get method without query", func(t *testing.T) {
		fetch := Fetch{}
		res, err := fetch.Get(GetParams{Client: client, Url: baseURL})
		
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.Status)
		assert.Equal(t, body, res.Data)
	})

	t.Run("Get method with query", func(t *testing.T) {		
		fetch := Fetch{}
		res, err := fetch.Get(GetParams{Client: client, Url: baseURL, Query: map[string]string{"foo": "bar"}})
		
		assert.Nil(t, err)
		assert.Equal(t, urlWithQuery, res.Url)
		assert.Equal(t, http.StatusOK, res.Status)
		assert.Equal(t, body, res.Data)
	})
}

func ExampleFetch_Get() {
	baseURL := "http://localhost:3000/ping"

	// Http mock to http://localhost:3000/ping. Always return "pong"
	gock.New(baseURL).Get("/ping").Reply(200).BodyString("pong")

	fetch := new(Fetch)
	client := &http.Client{}
	res, err := fetch.Get(GetParams{Client: client, Url: baseURL})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Staus code: %d, response body: %q", res.Status, res.Data)
	// Output: Staus code: 200, response body: "pong"
}