# gifs.com Golang Package


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

In this Go language package you'll be able to create basic gifs, and more complex functionality based on your application.
