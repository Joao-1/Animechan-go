package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"quotes/helpers"
)

const BaseURL = "http://animechan.melosh.space"
const CharacterPath = "/character"
const AnimePath = "/anime"
const QuotesOnlyPath = "/quotes"
const RandomPath = "/random"

type IParams interface {
	Anime(anime string) (Quote, error)
	Character(character string) (Quote, error)
	Only() (Quote, error)
}

type Params struct {
	anime string
	character string
}

type Quote struct {
	Anime string
	Character string
	Quote string
}

type QuoteAPIResponse struct {
	Key       int    `json:"key"`
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
	V       int    `json:"__v"`
}

type Random struct {
	Params
	fetch *helpers.Fetch
	client *http.Client
}

func (r *Random) Anime(anime string) (Quote, error) {
	path := BaseURL + RandomPath + AnimePath
	res, err := r.fetch.Get(helpers.GetParams{Client: r.client, Url: path, Query: map[string]string{"title": anime}})
	if err != nil {
		log.Fatal(err)
	}	

	var apiQuote QuoteAPIResponse
	errParse := json.Unmarshal([]byte(res.Data), &apiQuote)
	if errParse != nil {
		log.Fatal(err)
	}
	
	return Quote{Anime: apiQuote.Anime, Character: apiQuote.Character, Quote: apiQuote.Quote}, nil
}
func (r *Random) Character(character string) (Quote, error) {
	path := BaseURL + RandomPath + CharacterPath
	res, err := r.fetch.Get(helpers.GetParams{Client: r.client, Url: path, Query: map[string]string{"name": character}})
	if err != nil {
		log.Fatal(err)
	}	

	var apiQuote QuoteAPIResponse
	errParse := json.Unmarshal([]byte(res.Data), &apiQuote)
	if errParse != nil {
		log.Fatal(err)
	}
	
	return Quote{Anime: apiQuote.Anime, Character: apiQuote.Character, Quote: apiQuote.Quote}, nil
}

// Searches for a quote from a random anime and character
func (r *Random) Only() (Quote, error) {
	path := BaseURL + RandomPath

	res, err := r.fetch.Get(helpers.GetParams{Client: r.client, Url: path})
	if err != nil {
		log.Fatal(err)
	}	

	var apiQuote QuoteAPIResponse
	errParse := json.Unmarshal([]byte(res.Data), &apiQuote)
	if errParse != nil {
		log.Fatal(err)
	}

	return Quote{Anime: apiQuote.Anime, Character: apiQuote.Character, Quote: apiQuote.Quote}, nil
}

type Animechan struct {
	Client *http.Client
}

// Unlike the normal quote endpoint, it searches for a random one. It is possible to specify character or anime.
func (a *Animechan) Random() *Random {
	random := new(Random)
	fetch := new(helpers.Fetch)

	random.anime = AnimePath
	random.character = CharacterPath

	random.fetch = fetch
	random.client = a.Client
	
	return random
}

func main() {
	client := &http.Client{}

	animechan := Animechan{Client: client}
	quote, err := animechan.Random().Character("Kurama")
	if err != nil {
		panic(err)
	}

	fmt.Println(quote)
}