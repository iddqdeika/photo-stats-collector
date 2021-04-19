package definitions

import "io"

type ImageUrl string

type ImageStats struct {
	Height      int
	Width       int
	SizeInBytes int64
}

type UrlProvider interface {
	GetUrlList() []ImageUrl
}

type ImageDownloader interface {
	Download(url ImageUrl) (io.ReadCloser, error)
}

type StatCalculator interface {
	GetStats(data io.ReadCloser) (*ImageStats, error)
}
