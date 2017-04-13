package gifs

import (
	"os"
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

	if response.Files.Gif == "" {
		t.Fail()
	}

	if response.Embed == "" {
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
		Trim: &Trim{
			Start: 10,
			End:   20,
		},
		Safe: true,
	}

	response, err := input.Create()
	if err != nil {
		t.Fail()
	}

	if response.Files.Gif == "" {
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

	if file == "" {
		t.Fail()
	}

	t.Log("Downloaded File: ", file)

}

func TestUpload(t *testing.T) {

	input := &New{
		File:  "echo-hereweare.mp4",
		Title: "Echo Here We Are",
		Tags:  []string{"echo", "here", "we"},
	}

	response, err := input.Upload()
	if err != nil {
		t.Fail()
	}

	if response.Files.Gif == "" {
		t.Fail()
	}

	if response.Embed == "" {
		t.Fail()
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

	if gifFile == "" {
		t.Fail()
	}

	t.Log("Saved gif: ", gifFile)

}

func TestBulkUpload(t *testing.T) {

	files := []New{
		{
			File:  "echo-hereweare.mp4",
			Title: "New Video",
		},
		{
			File:  "echo-hereweare.mp4",
			Title: "New Video 2",
		},
		{
			File:  "echo-hereweare.mp4",
			Title: "New Video 3",
		},
	}

	bulk := Bulk{
		New: files,
	}

	response, err := bulk.Upload()
	if err != nil {
		t.Fail()
	}

	if len(response) != 3 {
		t.Fail()
	}

	for k, v := range response {
		t.Log("Upload", k, "- File Gif: ", v.Files.Gif)
	}

}
