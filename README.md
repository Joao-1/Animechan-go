
<h1 style="border-bottom: none" >Animechan + Golang</h1>
Animechan, as stated on its own website, is a free RESTful API that provides quality quotes. For greater ease of use of this API in multiple projects, this wrapper has been written.

Installation
============

To install, use `go get`:

    go get github.com/Joao-1/animechan-go
    
Examples
============

- Initialization
```go
	client := &http.Client{}
	animechan := Animechan{Client: client, BaseURL: "https://animechan.app/api/"}
```

------

- Random endpoint

Random quote
```go
	res, err := animechan.Random().Only()
	if err != nil { panic(err) }

	fmt.Printf("Anime: %q, Character: %q, Quote: %q", res.Anime, res.Character, res.Quote)
```

Random quote from a specific character
```go
	res, err := animechan.Random().Character("Madara Uchiha")
	if err != nil { panic(err) }

	fmt.Printf("Anime: %q, Character: %q, Quote: %q", res.Anime, res.Character, res.Quote)
```

Random quote from a specific anime
```go
	res, err := animechan.Random().Anime("Naruto")
	if err != nil { panic(err) }

	fmt.Printf("Anime: %q, Character: %q, Quote: %q", res.Anime, res.Character, res.Quote)
```
------

- Quotes endpoint

10 random quotes
```go
	res, err := animechan.Quotes().Only()
	if err != nil { panic(err) }

	fmt.Printf("all quotes: %v", res)
```
10 random quotes from the anime "Naruto"
```go
	page := 3
	res, err := animechan.Quotes().Anime("Naruto", &page)
	if err != nil { panic(err) }

	fmt.Printf("all quotes: %v", res)
```
10 random quotes from the character "Madara Uchiha"
```go
	page := 3
	res, err := animechan.Quotes().Character("Madara Uchiha", &page)
	if err != nil { panic(err) }

	fmt.Printf("all quotes: %v", res)
```
---
If you want to support me in any way:
<p align="center">
<br>
<a href="https://www.buymeacoffee.com/fukurou"><img src="https://github.com/appcraftstudio/buymeacoffee/raw/master/Images/snapshot-bmc-button.png" width="300"></a>
</p>
