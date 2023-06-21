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
	fetch *helpers.Fetch
	client *http.Client
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

type Quotes struct {
	fetch *helpers.Fetch
	client *http.Client
}

func (q *Quotes) Only() ([]Quote, error) {
	path := BaseURL + QuotesOnlyPath

	res, err := q.fetch.Get(helpers.GetParams{Client: q.client, Url: path})
	if err != nil {
		log.Fatal(err)
	}	

	var apiQuotes []QuoteAPIResponse
	errParse := json.Unmarshal([]byte(res.Data), &apiQuotes)
	if errParse != nil {
		log.Fatal(err)
	}

	var quotes []Quote
	for _, value := range apiQuotes {
		quotes = append(quotes, Quote{Anime: value.Anime, Character: value.Character, Quote: value.Quote})
	}

	return quotes, nil
}

func (q *Quotes) Anime(anime string, page *int) ([]Quote, error) {
	path := BaseURL + QuotesOnlyPath + AnimePath
	pageToSearch := 10

	if page != nil {pageToSearch = *page}

	res, err := q.fetch.Get(helpers.GetParams{Client: q.client, Url: path, Query: map[string]string{"title": anime, "page":fmt.Sprint(pageToSearch)}})
	if err != nil {
		log.Fatal(err)
	}	
	
	var apiQuotes []QuoteAPIResponse
	errParse := json.Unmarshal([]byte(res.Data), &apiQuotes)
	if errParse != nil {
		log.Fatal(err)
	}
	
	var quotes []Quote
	for _, value := range apiQuotes {
		quotes = append(quotes, Quote{Anime: value.Anime, Character: value.Character, Quote: value.Quote})
	}

	return quotes, nil
}

func (q *Quotes) Character(character string, page *int) ([]Quote, error) {
	path := BaseURL + QuotesOnlyPath + CharacterPath
	pageToSearch := 10

	if page != nil {pageToSearch = *page}

	res, err := q.fetch.Get(helpers.GetParams{Client: q.client, Url: path, Query: map[string]string{"name": character, "page":fmt.Sprint(pageToSearch)}})
	if err != nil {
		log.Fatal(err)
	}	

	var apiQuotes []QuoteAPIResponse
	errParse := json.Unmarshal([]byte(res.Data), &apiQuotes)
	if errParse != nil {
		log.Fatal(err)
	}

	var quotes []Quote
	for _, value := range apiQuotes {
		quotes = append(quotes, Quote{Anime: value.Anime, Character: value.Character, Quote: value.Quote})
	}

	return quotes, nil
}

type Animechan struct {
	Client *http.Client
}

// Unlike the normal quote endpoint, it searches for a random one. It is possible to specify character or anime.
func (a *Animechan) Random() *Random {
	random := new(Random)
	fetch := new(helpers.Fetch)

	random.fetch = fetch
	random.client = a.Client
	
	return random
}

func (a *Animechan) Quotes() *Quotes {
	quotes := new(Quotes)
	fetch := new(helpers.Fetch)

	quotes.fetch = fetch
	quotes.client = a.Client

	return quotes
}

func main() {
	client := &http.Client{}

	animechan := Animechan{Client: client}
	page := 1
	quote, err := animechan.Quotes().Character("Naruto", &page)
	if err != nil {
		panic(err)
	}

	fmt.Println(quote)
}