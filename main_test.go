package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestAnimechan(t *testing.T) {
	client := &http.Client{}
	
	animechan := Animechan{client}
	random := animechan.Random()

	anime := "Naruto"
	character := "Madara Uchiha"
	quote := "The longer you live... The more you realize that reality is just made of pain, suffering and emptiness..."
	body := fmt.Sprintf(`{"_id": "60393d7a234b061cfc607fb5","key": 3195,"anime": %q,"character": %q,"quote": %q,"__v": 0}`, anime, character, quote)
	// randomAnimeQuote := RandomPath + AnimePath + "?title=" + "Naruto"
	// gock.New(BaseURL).Get(RandomPath + AnimePath).Reply(200).JSON(map[string]string{"foo": "bar"})
	// gock.New(BaseURL).Get(RandomPath + CharacterPath).Reply(200).JSON(body)
	// gock.New(BaseURL).Get(QuotesOnlyPath).Reply(200).JSON(body)
	// gock.New(BaseURL).Get(QuotesOnlyPath + AnimePath).Reply(200).JSON(body)

	t.Run("Random Instance", func(t *testing.T) {
		assert.Equal(t, AnimePath, random.anime)
		assert.Equal(t, CharacterPath, random.character)
		assert.Implements(t, (*IParams)(nil), random)

		t.Run("Random - Only method", func(t *testing.T) {	
			defer gock.Off()
			gock.New(BaseURL).Get(RandomPath).Reply(200).JSON(body)

			res, err := random.Only()
	
			assert.Nil(t, err)
			assert.Equal(t, anime, res.Anime)
			assert.Equal(t, character, res.Character)
			assert.Equal(t, quote, res.Quote)
		})

		t.Run("Random - Anime method", func(t *testing.T) {
			defer gock.Off()
			gock.New(BaseURL).Get(RandomPath + AnimePath).Reply(200).JSON(body)
			
			res, err := random.Anime(anime)
	
			assert.Nil(t, err)
			
			assert.Equal(t, anime, res.Anime)
			assert.Equal(t, character, res.Character)
			assert.Equal(t, quote, res.Quote)
		})
		t.Run("Random - Character method", func(t *testing.T) {	
			defer gock.Off()
			gock.New(BaseURL).Get(RandomPath + CharacterPath).Reply(200).JSON(body)

			res, err := random.Character(character)
	
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