package gifs

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

func init() {
	Authentication = os.Getenv("AUTH")
}

func TestSimpleYoutube(t *testing.T) {

	file := &New{
		Source: "https://www.youtube.com/watch?v=V6wrI6DEZFk",
	}

	response, err := file.Create()
	if err != nil {
		t.Fail()
	}

	t.Log("Gif URL: ", response.Files.Gif)
	t.Log("Jpg URL: ", response.Files.Jpg)
	t.Log("Mp4 URL: ", response.Files.Mp4)
	t.Log("Page URL: ", response.Page)
	t.Log("oEmbed URL: ", response.Oembed)
	t.Log("Embed URL: ", response.Embed)

}

func TestYoutube(t *testing.T) {

	input := &New{
		Source: "https://www.youtube.com/watch?v=dDmQ0byhus4",
		Title:  "Cute Kitten Drinking From Sink",
		Tags:   []string{"cute", "kitten", "drinking"},
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

	t.Log("Gif URL: ", response.Files.Gif)
	t.Log("Jpg URL: ", response.Files.Jpg)
	t.Log("Mp4 URL: ", response.Files.Mp4)
	t.Log("Page URL: ", response.Page)
	t.Log("oEmbed URL: ", response.Oembed)
	t.Log("Embed URL: ", response.Embed)

}

func TestDownload(t *testing.T) {

	file := DownloadFile("echo-hereweare.mp4", "https://raw.githubusercontent.com/mediaelement/mediaelement-files/master/echo-hereweare.mp4")

	t.Log(file)

}

func TestUpload(t *testing.T) {

	input := &New{
		File:  "echo-hereweare.mp4",
		Title: "Echo Here We Are",
		Tags:  []string{"echo", "here", "we"},
	}

	response, err := input.Upload()
	if err != nil {
		panic(err)
	}

	t.Log("Gif URL: ", response.Files.Gif)
	t.Log("Jpg URL: ", response.Files.Jpg)
	t.Log("Mp4 URL: ", response.Files.Mp4)
	t.Log("Page URL: ", response.Page)
	t.Log("oEmbed URL: ", response.Oembed)
	t.Log("Embed URL: ", response.Embed)
}


func TestSaveGif(t *testing.T) {

	file := &New{
		Source: "https://www.youtube.com/watch?v=V6wrI6DEZFk",
	}

	response, err := file.Create()
	if err != nil {
		t.Fail()
	}

	gifFile := response.SaveGif()

	t.Log("Saved gif: ",gifFile)

}