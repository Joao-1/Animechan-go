package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"quotes/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAnimechan(t *testing.T) {
	client, mockClient := helpers.MockHttpClient()

	animechan := Animechan{client}
	random := animechan.Random()

	anime := "Naruto Shippūden"
	character := "Madara Uchiha"
	quote := "The longer you live... The more you realize that reality is just made of pain, suffering and emptiness..."
	body := fmt.Sprintf(`{"_id": "60393d7a234b061cfc607fb5","key": 3196,"anime": %q,"character": %q,"quote": %q,"__v": 0}`, anime, character, quote)
	
	setupMock := func() {
		mockClient.On("RoundTrip", mock.Anything).Return(&http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBufferString(body))}, nil)
	}

	restoreMock := func() {
		mockClient.ExpectedCalls = nil
	}
	t.Run("Random Instance", func(t *testing.T) {
		assert.Equal(t, AnimePath, random.anime)
		assert.Equal(t, CharacterPath, random.character)
		assert.Implements(t, (*IParams)(nil), random)

		t.Run("Random - Only method", func(t *testing.T) {
			setupMock()
			defer restoreMock()
	
			res, err := random.Only()
	
			assert.Nil(t, err)
			mockClient.AssertCalled(t, "NewRequest", "123")
			assert.Equal(t, anime, res.Anime)
			assert.Equal(t, character, res.Character)
			assert.Equal(t, quote, res.Quote)
		})
	
		t.Run("Random - Anime method", func(t *testing.T) {
			setupMock()
			defer restoreMock()
			
			res, err := random.Anime(anime)
	
			assert.Nil(t, err)
			
			assert.Equal(t, anime, res.Anime)
			assert.Equal(t, character, res.Character)
			assert.Equal(t, quote, res.Quote)
		})
		t.Run("Random - Character method", func(t *testing.T) {
			setupMock()
			defer restoreMock()
			
			res, err := random.Anime(anime)
	
			assert.Nil(t, err)
			assert.Equal(t, anime, res.Anime)
			assert.Equal(t, character, res.Character)
			assert.Equal(t, quote, res.Quote)
		})
	}) 
}

func ExampleRandom_Only() {
	client := &http.Client{}

	animechan := Animechan{Client: client}
	res, err := animechan.Random().Only()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Anime: %q, Character: %q, Quote: %q", res.Anime, res.Character, res.Quote)
}