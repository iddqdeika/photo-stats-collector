package main

import (
	"fmt"
	image_downloader "photo-stats-collector/pkg/image_downloader"
	"photo-stats-collector/pkg/stat_calculator"
)

func main() {
	dwdr := image_downloader.New()
	stc := stat_calculator.New()

	data, err := dwdr.Download("https://cdn.tom-tailor.com/img/1120_1490/1024691_27172_5.jpg")
	if err != nil {
		panic(err)
	}
	stat, err := stc.GetStats(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("stats: %v", stat)
}
