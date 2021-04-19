package definitions

import "io"

type ImageUrl string

type ImageStats struct {
	Url         string
	Height      int
	Width       int
	SizeInBytes int64
	Err         error
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

type TableWriter interface {
	WriteTable(data [][]string) error
}

type Processor interface {
	Process(url ImageUrl) ImageStats
}
