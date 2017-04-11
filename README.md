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
import "github.com/hunterlong/gifs"

func main() {

  input := &Import{
    Source: "https://www.youtube.com/watch?v=dDmQ0byhus4",
  }

  response := input.Create()

  fmt.Println("Gifs.com Gif URL: ",response.Files.Gif)
}
```
