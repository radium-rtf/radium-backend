package http

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Url  string
	Type string
	Size int64
}

func newYoutubeFileFromUrl(URL *url.URL) (*File, error) {
	v := URL.Query().Get("v")
	if v == "" {
		return nil, errors.New("empty video id")
	}
	if URL.Path != "/watch" {
		return nil, errors.New("url path can be /watch")
	}
	return &File{
		Url:  fmt.Sprintf("https://www.youtube.com/embed/%s", v),
		Name: "youtube",
		Type: "iframe",
		Size: 0,
	}, nil
}

func NewFileFromUrl(rawURL string) (*File, error) {
	URL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	if strings.HasSuffix(URL.Host, "youtube.com") {
		return newYoutubeFileFromUrl(URL)
	}
	response, err := http.Get(rawURL)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	contentType := response.Header.Get("Content-Type")
	contentLength := response.Header.Get("Content-Length")
	size, _ := strconv.ParseInt(contentLength, 10, 64)

	file := &File{Url: rawURL, Type: contentType, Size: size, Name: contentType}
	return file, nil
}
