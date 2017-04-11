# gifs.com Golang Package

<img width="400" align="right" src="https://j.gifs.com/r0G8wW.gif">

Golang is awesome, and so are gifs. Using gifs.com API you can create gifs from many many different sources.
Gifs.com API is already easy, but this simplifies it even more. In this Go language package you'll be able to create basic gifs, and more complex functionality based on your application.
<p></p>

##### Create a Gif from YouTube
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
