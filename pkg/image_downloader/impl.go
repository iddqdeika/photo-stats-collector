package image_downloader

import (
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
	return resp.Body, nil
}
