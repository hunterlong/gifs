# gifs.com Golang Package

<img width="400" align="right" src="https://j.gifs.com/r0G8wW.gif">

[![Build Status](https://travis-ci.org/hunterlong/gifs.svg?branch=master)](https://travis-ci.org/hunterlong/gifs) [![Coverage Status](https://coveralls.io/repos/github/hunterlong/gifs/badge.svg?branch=master)](https://coveralls.io/github/hunterlong/gifs?branch=master) [![GoDoc](https://godoc.org/github.com/hunterlong/gifs?status.svg)](https://godoc.org/github.com/hunterlong/gifs) [![Go Report Card](https://goreportcard.com/badge/github.com/hunterlong/gifs)](https://goreportcard.com/report/github.com/hunterlong/gifs)

Golang is awesome, and so are gifs. Using gifs.com API you can create gifs from many many different sources.
Gifs.com API is already easy, but this simplifies it even more. You'll be able to create basic gifs, and more complex functionality based on your application.

:white_check_mark: Create gifs via URL

:white_check_mark: Upload file and create gif

<p></p>

# Install
```go
go get -u github.com/hunterlong/gifs
```

```go
import "github.com/hunterlong/gifs"
```

# Simple Usage
```go
input := &gifs.New{
  Source: "https://www.youtube.com/watch?v=dDmQ0byhus4",
}

response := input.Create()

fmt.Println("Gifs.com Gif URL: ",response.Files.Gif)
```

# Advanced Usage
```go
input := &gifs.New{
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
if err != nil {
    panic(err)
}

fmt.Println("Gif URL: ", response.Files.Gif)
fmt.Println("Jpg URL: ", response.Files.Jpg)
fmt.Println("Mp4 URL: ", response.Files.Mp4)
fmt.Println("Page URL: ", response.Page)
fmt.Println("oEmbed URL: ", response.Oembed)
fmt.Println("Embed URL: ", response.Embed)
```

# Saving gif
```go
file := &gifs.New{
    Source: "https://www.youtube.com/watch?v=V6wrI6DEZFk",
}

response, _ := file.Create()

gifFile := response.SaveGif()

fmt.Println("Saved gif: ", gifFile)

```

# Upload MP4
```go
input := &gifs.New{
    File:  "video.mp4",
    Title: "Nice Video Title",
    Tags:  []string{"gifs", "are", "awesome"},
}

response, err := input.Upload()
if err != nil {
    panic(err)
}

fmt.Println("Gif URL: ", response.Files.Gif)
```

# Bulk Upload
```go
// work in progress!
```