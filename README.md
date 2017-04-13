# [gifs.com](https://gifs.com) Golang Package

<img width="160" align="right" src="https://j.gifs.com/wjN768.gif">

[![Build Status](https://travis-ci.org/hunterlong/gifs.svg?branch=master)](https://travis-ci.org/hunterlong/gifs) [![Coverage Status](https://coveralls.io/repos/github/hunterlong/gifs/badge.svg?branch=master)](https://coveralls.io/github/hunterlong/gifs?branch=master) [![GoDoc](https://godoc.org/github.com/hunterlong/gifs?status.svg)](https://godoc.org/github.com/hunterlong/gifs) [![Go Report Card](https://goreportcard.com/badge/github.com/hunterlong/gifs)](https://goreportcard.com/report/github.com/hunterlong/gifs)

Golang is awesome, and so are gifs. Using [gifs.com API](http://docs.gifs.com/docs/getting-started) you can create gifs from many many different sources.
Gifs.com API is already easy, but this simplifies it even more. You'll be able to create basic gifs, and more complex functionality based on your application.

:white_check_mark: Create gifs via URL (youtube, vimeo, remote mp4, etc)

:white_check_mark: Upload mp4/image & Download gif

:white_check_mark: Bulk Upload

:white_check_mark: Trim gif (start/end in seconds)

Gifs.com allows you to send any instagram, twitter, facebook, vine, .gif, .mp4, .webm, giphy, imgur, gfycat, or streamable links. :thumbsup:

<p></p>

# Install
```go
go get -u github.com/hunterlong/gifs
```

```go
import "github.com/hunterlong/gifs"
```

# Simple Usage
Be :sunglasses: and use API keys for added features and limits. [gifs.com Authentication](http://docs.gifs.com/docs/authentication-key)  API does *not* require keys.
```go
gifs.Authentication = "gifs00YOURkey2929"
```

Let's make a gif from a YouTube video.
```go
input := &gifs.New{
  Source: "https://www.youtube.com/watch?v=dDmQ0byhus4",
}

response := input.Create()

fmt.Println("Gifs.com Gif URL: ",response.Files.Gif)
```

# Advanced Usage
You can add all the gifs.com attributes and even tag the gif with a user. You can trim your gif by inserting start and end in seconds. There's even a 'safe' option.
```go
input := &gifs.New{
  Source: "https://www.youtube.com/watch?v=dDmQ0byhus4",
  Title: "Cute Kitten Drinking From Sink",
  Tags: []string{"cute","kitten","drinking"},
  Attribution: &Attribution{
    Site: "twitter",
    User: "stronghold2d",
  },
  Trim: &Trim{
    Start: 10,
    End:   20,
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
You've seen the responses from gifs.com API, but now I'd like to download the gif and save it locally.
<img width="300" align="right" src="https://j.gifs.com/2RpGgv.gif">
```go
file := &gifs.New{
    Source: "https://www.youtube.com/watch?v=V6wrI6DEZFk",
}

response, _ := file.Create()

gifFile := response.SaveGif()

fmt.Println("Saved gif: ", gifFile)
```

# Upload MP4
Oh whaaat. I have a mp4 file named 'video.mp4' in my current directory and I'd love to have it as a gif.
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
Now I've got the hang of it, I don't need to go one-by-one. Include an array of New files, and Upload! The response will give you an array for each uploaded file.
<img width="300" align="right" src="https://j.gifs.com/nZAYM5.gif">

```go
// work in progress!
files := []gifs.New{
    {
        File:  "video1.mp4",
        Title: "New Video",
    },
    {
        File:  "video2.mp4",
        Title: "New Video 2",
    },
    {
        File:  "video3.mp4",
        Title: "New Video 3",
    },
}

bulk := gifs.Bulk{
    New: files,
}

response, err := bulk.Upload()
if err != nil {
    panic(err)
}

for k, v := range response {
    fmt.Println("File #", k, " Uploaded. File Gif: ", v.Files.Gif)
}

fmt.Println("Uploaded", len(response), "Files")
```

# License
Fun project :boom: How often do you see gifs on github!? Released MIT license with testing.
