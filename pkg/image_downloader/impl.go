package image_downloader

import (
	"fmt"
	"io"
	"net/http"
	"photo-stats-collector/definitions"
)

func New() definitions.ImageDownloader {
	return &imageDownloader{}
}

type imageDownloader struct {
}

func (i *imageDownloader) Download(url definitions.ImageUrl) (io.ReadCloser, error) {
	resp, err := http.Get(string(url))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("cant download, status code is %v", resp.StatusCode)
	}
	return resp.Body, nil
}
