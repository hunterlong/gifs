# gifs.com Golang Package

<img width="400" align="right" src="https://j.gifs.com/r0G8wW.gif">

Golang is awesome, and so are gifs. Using gifs.com API you can create gifs from many many different sources.
Gifs.com API is already easy, but this simplifies it even more. You'll be able to create basic gifs, and more complex functionality based on your application.

:white_check_mark: Create gifs via URL

:white_check_mark: Upload file and create gif

:white_check_mark: Bulk Upload and create gifs

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
