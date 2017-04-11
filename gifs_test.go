package gifs

import (
	"testing"
	"net/url"
	"strings"
	"os"
	"io"
	"net/http"
)


func init() {
	Authentication = os.Getenv("AUTH")
}


func TestSimpleYoutube(t *testing.T) {

	input := &Import{
		Source: "https://www.youtube.com/watch?v=e9GZ_GD4R4s",
	}

	response, err := input.Create()
	if err != nil {
		t.Fail()
	}

	t.Log("Gif URL: ",response.Files.Gif)
	t.Log("Jpg URL: ",response.Files.Jpg)
	t.Log("Mp4 URL: ",response.Files.Mp4)
	t.Log("Page URL: ",response.Page)
	t.Log("oEmbed URL: ",response.Oembed)
	t.Log("Embed URL: ",response.Embed)

}


func TestYoutube(t *testing.T) {

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
	if err != nil {
		t.Fail()
	}

	t.Log("Gif URL: ",response.Files.Gif)
	t.Log("Jpg URL: ",response.Files.Jpg)
	t.Log("Mp4 URL: ",response.Files.Mp4)
	t.Log("Page URL: ",response.Page)
	t.Log("oEmbed URL: ",response.Oembed)
	t.Log("Embed URL: ",response.Embed)

}



func TestDownload(t *testing.T) {

	file := DownloadFile("echo-hereweare.mp4", "https://raw.githubusercontent.com/mediaelement/mediaelement-files/master/echo-hereweare.mp4")

	t.Log(file)

}


func TestUpload(t *testing.T) {

	input := &Import{
		File: "echo-hereweare.mp4",
		Title: "Echo Here We Are",
		Tags: []string{"echo","here","we"},
	}

	response, err := input.Upload()
	if err != nil {
		panic(err)
	}

	t.Log("Gif URL: ",response.Files.Gif)
	t.Log("Jpg URL: ",response.Files.Jpg)
	t.Log("Mp4 URL: ",response.Files.Mp4)
	t.Log("Page URL: ",response.Page)
	t.Log("oEmbed URL: ",response.Oembed)
	t.Log("Embed URL: ",response.Embed)
}



func DownloadFile(n string, rawURL string) string {

	fileURL, err := url.Parse(rawURL)

	if err != nil {
		panic(err)
	}

	path := fileURL.Path

	segments := strings.Split(path, "/")

	fileName := segments[4]

	file, err := os.Create(n)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	check := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := check.Get(rawURL) // add a filter to check redirect

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.Copy(file, resp.Body)

	if err != nil {
		panic(err)
	}
	return fileName
}