package helpers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T){
	baseURL := "https://localhost:3000"

	t.Run("Get method", func(t *testing.T) {
		client, mockClient := MockHttpClient()
		
		body := "Hello, World"
		mockClient.On("RoundTrip", mock.Anything).Return(&http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString(body))}, nil)

		fetch := Fetch{}
		res, err := fetch.Get(baseURL, client)
		
		fmt.Println(res)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.Status)
		assert.Equal(t, body, res.Data)
	})
}

func ExampleFetch_Get() {
	Server()

	fetch := new(Fetch)
	client := &http.Client{}
	res, err := fetch.Get("http://localhost:3000/ping", client)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Staus code: %d, response body: %q", res.Status, res.Data)
	// Output: Staus code: 200, response body: "\"pong\"\n"
}