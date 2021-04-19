package url_provider

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"photo-stats-collector/definitions"
)

func NewXlsx(filename string, sheetname string) (definitions.UrlProvider, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error while opening file %v, error: %v", filename, err)
	}
	rows, err := f.GetRows(sheetname)
	if err != nil {
		return nil, fmt.Errorf("error while reading sheet data, error: %v", err)
	}
	urls := make([]definitions.ImageUrl, len(rows))
	for i, row := range rows {
		var val definitions.ImageUrl
		if len(row) > 0 {
			val = definitions.ImageUrl(row[0])
		}
		urls[i] = val
	}
	return &urlProvider{urls: urls}, nil
}

type urlProvider struct {
	urls []definitions.ImageUrl
}

func (u *urlProvider) GetUrlList() []definitions.ImageUrl {
	return u.urls
}
