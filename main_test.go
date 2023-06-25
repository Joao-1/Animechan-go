package main

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

var anime = "Naruto"
var character = "Madara Uchiha"
var quote = "The longer you live... The more you realize that reality is just made of pain, suffering and emptiness..."
var singleQuote = fmt.Sprintf(`{"_id": "60393d7a234b061cfc607fb5","key": 3195,"anime": %q,"character": %q,"quote": %q,"__v": 0}`, anime, character, quote)

func TestAnimechan(t *testing.T) {
	client := &http.Client{}
	
	animechan := Animechan{client}
	t.Run("Quotes Instance", func(t *testing.T) {
		assert.Implements(t, (*IQuotes)(nil), animechan.Quotes())

		t.Run("Quotes - Only method", func(t *testing.T) {	
			defer gock.Off()
			gock.New(BaseURL).Get(QuotesOnlyPath).Reply(200).JSON("[" + strings.Join([]string{singleQuote, singleQuote, singleQuote, singleQuote, singleQuote, singleQuote, singleQuote, singleQuote, singleQuote, singleQuote}, ",") + "]")
			
			res, err := animechan.Quotes().Only()

			assert.Nil(t, err)
			assert.Len(t, res, 10)
			
			for _, q := range res {
				assert.Equal(t, anime, q.Anime)
				assert.Equal(t, character, q.Character)
				assert.Equal(t, quote, q.Quote)
			}
		})

		t.Run("Quotes - Anime method", func(t *testing.T) {
			defer gock.Off()
			gock.New(BaseURL).Get(QuotesOnlyPath + AnimePath).Reply(200).JSON("[" + strings.Join([]string{singleQuote, singleQuote, singleQuote, singleQuote, singleQuote, singleQuote, singleQuote, singleQuote, singleQuote, singleQuote}, ",") + "]")
			
			res, err := animechan.Quotes().Anime(anime, nil)
	
			assert.Nil(t, err)
			assert.Len(t, res, 10)
			
			for _, q := range res {
				assert.Equal(t, anime, q.Anime)
				assert.Equal(t, character, q.Character)
				assert.Equal(t, quote, q.Quote)
			}
		})

		t.Run("Quotes - Anime method - pagination", func(t *testing.T) {
			defer gock.Off()
			gock.New(BaseURL).Get(QuotesOnlyPath + AnimePath).Reply(200).JSON("[" + strings.Join([]string{singleQuote, singleQuote, singleQuote}, ",") + "]")

			page := 3
			res, err := animechan.Quotes().Anime(anime, &page)
	
			assert.Nil(t, err)
			assert.Len(t, res, page)
			
			for _, q := range res {
				assert.Equal(t, anime, q.Anime)
				assert.Equal(t, character, q.Character)
				assert.Equal(t, quote, q.Quote)
			}
		})
		t.Run("Quotes - Character method", func(t *testing.T) {	
			defer gock.Off()
			gock.New(BaseURL).Get(QuotesOnlyPath + CharacterPath).Reply(200).JSON("[" + strings.Join([]string{singleQuote, singleQuote, singleQuote, singleQuote, singleQuote, singleQuote, singleQuote, singleQuote, singleQuote, singleQuote}, ",") + "]")

			res, err := animechan.Quotes().Character(anime, nil)
	
			assert.Nil(t, err)
			assert.Len(t, res, 10)
			
			for _, q := range res {
				assert.Equal(t, anime, q.Anime)
				assert.Equal(t, character, q.Character)
				assert.Equal(t, quote, q.Quote)
			}
		})

		t.Run("Quotes - Character method - Pagination", func(t *testing.T) {	
			defer gock.Off()
			gock.New(BaseURL).Get(QuotesOnlyPath + CharacterPath).Reply(200).JSON("[" + strings.Join([]string{singleQuote, singleQuote, singleQuote}, ",") + "]")

			page := 3
			res, err := animechan.Quotes().Character(anime, &page)
	
			assert.Nil(t, err)
			assert.Len(t, res, page)
			
			for _, q := range res {
				assert.Equal(t, anime, q.Anime)
				assert.Equal(t, character, q.Character)
				assert.Equal(t, quote, q.Quote)
			}
		})
	})

	t.Run("Random Instance", func(t *testing.T) {
		assert.Implements(t, (*IRandom)(nil), animechan.Random())
		t.Run("Random - Only method", func(t *testing.T) {	
			defer gock.Off()
			gock.New(BaseURL).Get(RandomPath).Reply(200).JSON(singleQuote)

			res, err := animechan.Random().Only()
	
			assert.Nil(t, err)
			assert.Equal(t, anime, res.Anime)
			assert.Equal(t, character, res.Character)
			assert.Equal(t, quote, res.Quote)
		})

		t.Run("Random - Anime method", func(t *testing.T) {
			defer gock.Off()
			gock.New(BaseURL).Get(RandomPath + AnimePath).Reply(200).JSON(singleQuote)
			
			res, err := animechan.Random().Anime(anime)
	
			assert.Nil(t, err)
			
			assert.Equal(t, anime, res.Anime)
			assert.Equal(t, character, res.Character)
			assert.Equal(t, quote, res.Quote)
		})
		t.Run("Random - Character method", func(t *testing.T) {	
			defer gock.Off()
			gock.New(BaseURL).Get(RandomPath + CharacterPath).Reply(200).JSON(singleQuote)

			res, err := animechan.Random().Character(character)
	
			assert.Nil(t, err)
			assert.Equal(t, anime, res.Anime)
			assert.Equal(t, character, res.Character)
			assert.Equal(t, quote, res.Quote)
		})
	})
}

func TestFormatOneQuote(t *testing.T) {
	t.Run("With data", func(t *testing.T) {
		formattedQuote, err := formatOneQuote(singleQuote) 

		assert.Nil(t, err)
		assert.Equal(t, anime, formattedQuote.Anime)
		assert.Equal(t, character, formattedQuote.Character)
		assert.Equal(t, quote, formattedQuote.Quote)
	})

	t.Run("Without data", func(t *testing.T) {
		formattedQuote, err := formatOneQuote("{}") 

		assert.Nil(t, err)
		assert.Empty(t, formattedQuote)
	})
}

func TestFormatManyQuotes(t *testing.T) {
	formattedQuote, err := formatManyQuote("[" + strings.Join([]string{singleQuote, singleQuote, singleQuote}, ",") + "]") 

	assert.Nil(t, err)
	assert.Len(t, formattedQuote, 3)
	
	for _, q := range formattedQuote {
		assert.Equal(t, anime, q.Anime)
		assert.Equal(t, character, q.Character)
		assert.Equal(t, quote, q.Quote)
	}

	t.Run("Without data", func(t *testing.T) {
		formattedQuote, err := formatManyQuote("[]") 

		assert.Nil(t, err)
		assert.Empty(t, len(formattedQuote))
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

func ExampleRandom_Anime(){
	client := &http.Client{}

	animechan := Animechan{Client: client}
	res, err := animechan.Random().Anime("Naruto")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Anime: %q, Character: %q, Quote: %q", res.Anime, res.Character, res.Quote)
}

func ExampleRandom_Character(){
	client := &http.Client{}

	animechan := Animechan{Client: client}
	res, err := animechan.Random().Character("Madara Uchiha")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Anime: %q, Character: %q, Quote: %q", res.Anime, res.Character, res.Quote)
}

func ExampleQuotes_Only() {
	client := &http.Client{}

	animechan := Animechan{Client: client}
	res, err := animechan.Quotes().Only()
	if err != nil {
		panic(err)
	}

	fmt.Printf("all quotes: %v", res)
}

func ExampleQuotes_Anime() {
	client := &http.Client{}

	animechan := Animechan{Client: client}
	page := 3
	res, err := animechan.Quotes().Anime("Naruto", &page)
	if err != nil {
		panic(err)
	}

	fmt.Printf("all quotes: %v", res)
}

func ExampleQuotes_Character() {
	client := &http.Client{}

	animechan := Animechan{Client: client}
	page := 3
	res, err := animechan.Quotes().Character("Madara Uchiha", &page)
	if err != nil {
		panic(err)
	}

	fmt.Printf("all quotes: %v", res)
}