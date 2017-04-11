# gifs.com Golang Package

<img width="400" align="right" src="https://j.gifs.com/r0G8wW.gif">

[![Build Status](https://travis-ci.org/hunterlong/gifs.svg?branch=master)](https://travis-ci.org/hunterlong/gifs) [![Coverage Status](https://coveralls.io/repos/github/hunterlong/gifs/badge.svg?branch=master)](https://coveralls.io/github/hunterlong/gifs?branch=master) [![GoDoc](https://godoc.org/github.com/hunterlong/gifs?status.svg)](https://godoc.org/github.com/hunterlong/gifs) [![Go Report Card](https://goreportcard.com/badge/github.com/hunterlong/gifs)](https://goreportcard.com/report/github.com/hunterlong/gifs)

Golang is awesome, and so are gifs. Using gifs.com API you can create gifs from many many different sources.
Gifs.com API is already easy, but this simplifies it even more. You'll be able to create basic gifs, and more complex functionality based on your application.

:white_check_mark: Create gifs via URL

:white_check_mark: Upload file and create gif

<p></p>

# Simple Usage
```go
input := &Import{
  Source: "https://www.youtube.com/watch?v=dDmQ0byhus4",
}

response := input.Create()

fmt.Println("Gifs.com Gif URL: ",response.Files.Gif)
```

# Advanced Usage
```go
input := &Import{
  Source: "https://www.youtube.com/watch?v=dDmQ0byhus4",
  Title: "Cute Kitten Drinking From Sink",
  Tags: []string{"cute","kitten","drinking"},
  Attribution: &Attribution{
    Site: "twitter",
    User: "stronghold2d",
  },
  Safe: true,
}

response, err := input.Create()
```
