package gifs

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"mime/multipart"
	"path/filepath"
	"io"
	"os"
)

const ApiEndpoint string = "https://api.gifs.com"
var Authentication string


type Import struct {
	Source string `json:"source,omitempty"`
	File string `json:"-"`
	Title string `json:"title,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Attribution *Attribution `json:"attribution,omitempty"`
	Safe bool `json:"nsfw,omitempty"`
}

type Attribution struct {
	Site string `json:"site,omitempty"`
	User string `json:"user,omitempty"`
	Url string `json:"url,omitempty"`
}


type Success struct {
	Response ImportResponse `json:"success"`
}

type ImportResponse struct {
	Page string `json:"page"`
	Files struct {
		     Gif string `json:"gif"`
		     Jpg string `json:"jpg"`
		     Mp4 string `json:"mp4"`
		     Webm string `json:"webm"`
	     } `json:"files"`
	Oembed string `json:"oembed"`
	Embed string `json:"embed"`
	Meta struct {
		     Duration string `json:"duration"`
		     Height string `json:"height"`
		     Width string `json:"width"`
	     } `json:"meta"`
}


func (r *ImportResponse) Gif() string {
	return r.Files.Gif
}

func (r *ImportResponse) Jpg() string {
	return r.Files.Jpg
}

func (r *ImportResponse) Mp4() string {
	return r.Files.Mp4
}

func checkErr(e error){

}


func (i *Import) Create() (*ImportResponse, error)  {
	var err error
	req, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(req))
	res, err := SendRequest(req, "/media/import")
	if err != nil {
		return nil, err
	}
	var d Success
	json.Unmarshal(res,&d)
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
		fmt.Println("Could not connect to server at: ",ApiEndpoint)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	fmt.Println(string(body))
	return body, err
}



func (i *Import) Upload() (*ImportResponse, error)  {
	var err error
	res, err := UploadRequest(i, i.File)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	var d Success
	json.Unmarshal(res,&d)
	return &d.Response, err

}

func UploadRequest(i *Import, fileName string) ([]byte, error) {
	path, _ := os.Getwd()
	path += "/"+fileName
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
		panic(err)
	} else {
		body = &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			panic(err)
		}
		resp.Body.Close()
	}

	return body.Bytes(), err
}
