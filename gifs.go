package gifs

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const ApiEndpoint string = "https://api.gifs.com"

var Authentication string

type New struct {
	Source      string       `json:"source,omitempty"`
	File        string       `json:"-"`
	Title       string       `json:"title,omitempty"`
	Tags        []string     `json:"tags,omitempty"`
	Attribution *Attribution `json:"attribution,omitempty"`
	Safe        bool         `json:"nsfw,omitempty"`
}

type Attribution struct {
	Site string `json:"site,omitempty"`
	User string `json:"user,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Success struct {
	Response ImportResponse `json:"success"`
}

type ImportResponse struct {
	Page  string `json:"page"`
	Files struct {
		Gif  string `json:"gif"`
		Jpg  string `json:"jpg"`
		Mp4  string `json:"mp4"`
		Webm string `json:"webm"`
	} `json:"files"`
	Oembed string `json:"oembed"`
	Embed  string `json:"embed"`
	Meta   struct {
		Duration string `json:"duration"`
		Height   string `json:"height"`
		Width    string `json:"width"`
	} `json:"meta"`
}

func (r *ImportResponse) SaveGif() string {
	file := DownloadFile("newgif.gif", r.Files.Gif)
	return file
}

func (i *New) Create() (*ImportResponse, error) {
	var err error
	req, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	res, err := SendRequest(req, "/media/import")
	if err != nil {
		return nil, err
	}
	var d Success
	json.Unmarshal(res, &d)
	return &d.Response, err
}

func SendRequest(input []byte, method string) ([]byte, error) {
	req, err := http.NewRequest("POST", ApiEndpoint+method, bytes.NewBuffer(input))
	if Authentication != "" {
		req.Header.Set("Gifs-API-Key", Authentication)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Could not connect to server at: ", ApiEndpoint)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Could not read response from API ")
		return nil, err
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	return body, err
}

func (i *New) Upload() (*ImportResponse, error) {
	var err error
	res, err := UploadRequest(i, i.File)
	if err != nil {
		return nil, err
	}
	var d Success
	json.Unmarshal(res, &d)
	return &d.Response, err

}

func UploadRequest(i *New, fileName string) ([]byte, error) {
	path, _ := os.Getwd()
	path += "/" + fileName
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	if i.Title != "" {
		writer.WriteField("title", i.Title)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", ApiEndpoint+"/media/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else {
		body = &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			return nil, err
		}
		resp.Body.Close()
	}

	return body.Bytes(), err
}

func DownloadFile(n string, rawURL string) string {

	file, err := os.Create(n)

	if err != nil {
		return nil
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
		return nil
	}
	return n
}
