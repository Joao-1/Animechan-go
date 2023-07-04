package animechan

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joao-1/animechan-go/helpers"
)

const BaseURL = "https://animechan.xyz/api"
const CharacterPath = "/character"
const AnimePath = "/anime"
const QuotesOnlyPath = "/quotes"
const RandomPath = "/random"

type IRandom interface {
	Anime(anime string) (Quote, error)
	Character(character string) (Quote, error)
	Only() (Quote, error)
}

type IQuotes interface {
	Anime(anime string, page *int) ([]Quote, error)
	Character(character string, page *int) ([]Quote, error)
	Only() ([]Quote, error)
}

// Quote struct.
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

// Quotes is the struct for the quotes endpoint.
type Random struct {
	fetch *helpers.Fetch
	client *http.Client
}

// Searches for a quote from a random anime and character
func (r *Random) Only() (Quote, error) {
	path := BaseURL + RandomPath

	res, err := r.fetch.Get(helpers.GetParams{Client: r.client, Url: path})
	if err != nil { return Quote{}, err }

	return formatOneQuote(res.Data)
}

// Searches for a random quote from a specific anime
func (r *Random) Anime(anime string) (Quote, error) {
	path := BaseURL + RandomPath + AnimePath

	res, err := r.fetch.Get(helpers.GetParams{Client: r.client, Url: path, Query: map[string]string{"title": anime}})
	if err != nil { return Quote{}, err }

	return formatOneQuote(res.Data)
}

// Search for a random quote from a specific character
func (r *Random) Character(character string) (Quote, error) {
	path := BaseURL + RandomPath + CharacterPath

	res, err := r.fetch.Get(helpers.GetParams{Client: r.client, Url: path, Query: map[string]string{"name": character}})
	if err != nil { return Quote{}, err }

	return formatOneQuote(res.Data)
}

// Quotes is the struct for the quotes endpoint. 
type Quotes struct {
	fetch *helpers.Fetch
	client *http.Client
}

// Searches for 10 quotes from random anime and character.
func (q *Quotes) Only() ([]Quote, error) {
	path := BaseURL + QuotesOnlyPath

	res, err := q.fetch.Get(helpers.GetParams{Client: q.client, Url: path})
	if err != nil { return []Quote{}, err }	

	return formatManyQuote(res.Data)
}

// Searches for quotes from a specific anime. It is possible to specify page.
func (q *Quotes) Anime(anime string, page *int) ([]Quote, error) {
	path := BaseURL + QuotesOnlyPath + AnimePath
	pageToSearch := 10

	if page != nil {pageToSearch = *page}

	res, err := q.fetch.Get(helpers.GetParams{Client: q.client, Url: path, Query: map[string]string{"title": anime, "page":fmt.Sprint(pageToSearch)}})
	if err != nil { return []Quote{}, err }		
	
	return formatManyQuote(res.Data)
}
// Searches for quotes from a specific character. It is possible to specify page.
func (q *Quotes) Character(character string, page *int) ([]Quote, error) {
	path := BaseURL + QuotesOnlyPath + CharacterPath
	pageToSearch := 10

	if page != nil {pageToSearch = *page}

	res, err := q.fetch.Get(helpers.GetParams{Client: q.client, Url: path, Query: map[string]string{"name": character, "page":fmt.Sprint(pageToSearch)}})
	if err != nil { return []Quote{}, err }	

	return formatManyQuote(res.Data)
}

// Animechan is the main struct for the package. It contains the client for the http requests.
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

// Searches for quotes from a specific anime or character. It is possible to specify page.
func (a *Animechan) Quotes() *Quotes {
	quotes := new(Quotes)
	fetch := new(helpers.Fetch)

	quotes.fetch = fetch
	quotes.client = a.Client

	return quotes
}

func formatOneQuote(data string) (Quote, error) {
	var apiQuote QuoteAPIResponse
	
	errParse := json.Unmarshal([]byte(data), &apiQuote)
	if errParse != nil { return Quote{}, errParse }	

	return Quote{Anime: apiQuote.Anime, Character: apiQuote.Character, Quote: apiQuote.Quote}, nil
}

func formatManyQuote(data string) ([]Quote, error) {
	var apiQuotes []QuoteAPIResponse
	errParse := json.Unmarshal([]byte(data), &apiQuotes)
	if errParse != nil { return []Quote{}, errParse }	

	var quotes []Quote
	for _, value := range apiQuotes {
		quotes = append(quotes, Quote{Anime: value.Anime, Character: value.Character, Quote: value.Quote})
	}

	return quotes, nil
}
