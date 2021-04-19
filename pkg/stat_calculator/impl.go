package stat_calculator

import (
	"bytes"
	_ "golang.org/x/image/webp"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"photo-stats-collector/definitions"
)

func New() definitions.StatCalculator {
	return &statCalculator{}
}

type statCalculator struct {
}

func (s statCalculator) GetStats(data io.ReadCloser) (*definitions.ImageStats, error) {
	defer data.Close()
	rw := bytes.NewBuffer(nil)
	size, err := io.Copy(rw, data)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(rw)
	if err != nil {
		return nil, err
	}
	return &definitions.ImageStats{
		Height:      img.Bounds().Dy(),
		Width:       img.Bounds().Dx(),
		SizeInBytes: size,
	}, nil
}
