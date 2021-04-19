package root

import (
	"context"
	"fmt"
	"github.com/iddqdeika/rrr"
	"github.com/iddqdeika/rrr/helpful"
	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
	"photo-stats-collector/definitions"
	"photo-stats-collector/pkg/image_downloader"
	"photo-stats-collector/pkg/processor"
	"photo-stats-collector/pkg/stat_calculator"
	"photo-stats-collector/pkg/table_writer"
	url_provider "photo-stats-collector/pkg/url-provider"
	"strconv"
	"sync"
	"time"
)

const (
	defaultRequestIntervalInMillis = 300
)

func New() rrr.Root {
	return &root{}
}

type root struct {
	l helpful.Logger

	ri time.Duration

	up definitions.UrlProvider
	tw definitions.TableWriter
	p  definitions.Processor
}

func (r *root) Register() []error {
	// логгер
	r.l = helpful.DefaultLogger.WithLevel(helpful.LogInfo)
	r.l.Infof("logger initialized, proceeding register")
	defer r.l.Infof("register finished")

	cfg, err := helpful.NewJsonCfg("config.json")
	if err != nil {
		return []error{errors.Wrap(err, "cant init config")}
	}

	e := func(err error) []error {
		return []error{err}
	}
	filename, err := cfg.GetString("excel_file_name")
	if err != nil {
		return e(err)
	}

	inputSheet, err := cfg.GetString("input_sheet_name")
	if err != nil {
		return e(err)
	}

	outputSheetName, err := cfg.GetString("output_sheet_name")
	if err != nil {
		return e(err)
	}
	requestInterval, err := cfg.GetInt("request_interval_in_millis")
	if err != nil {
		requestInterval = defaultRequestIntervalInMillis
	}
	r.ri = time.Millisecond * time.Duration(requestInterval)

	r.up, err = url_provider.NewXlsx(filename, inputSheet)
	if err != nil {
		return e(err)
	}

	r.p = processor.New(image_downloader.New(), stat_calculator.New())
	r.tw = table_writer.New(filename, outputSheetName)

	return nil
}

func (r *root) Resolve(ctx context.Context) error {
	urls := r.up.GetUrlList()
	balancer := make(chan struct{}, 10)
	ticker := time.NewTicker(r.ri)
	defer ticker.Stop()
	wg := sync.WaitGroup{}

	pb := progressbar.Default(int64(len(urls)))
	defer pb.Finish()

	stats := make([]definitions.ImageStats, len(urls))
	m := &sync.Mutex{}
	for i, url := range urls {
		balancer <- struct{}{}
		<-ticker.C
		wg.Add(1)
		go func(url definitions.ImageUrl, i int) {
			stat := r.p.Process(url)
			m.Lock()
			stats[i] = stat
			m.Unlock()
			wg.Done()
			<-balancer
			pb.Add(1)
		}(url, i)
	}

	wg.Wait()
	rows := make([][]string, len(stats))
	for i, stat := range stats {
		row := make([]string, 4)
		row[0] = stat.Url
		if stat.Err != nil {
			row[1] = stat.Err.Error()
		} else {
			row[1] = strconv.Itoa(int(stat.SizeInBytes)) + " bytes"
			row[2] = strconv.Itoa(stat.Width) + "x" + strconv.Itoa(stat.Height)
		}
		rows[i] = row
	}

	return r.tw.WriteTable(rows)
}

func (r *root) Release() error {
	fmt.Println("PRESS ENTER")
	fmt.Scanln()
	return nil
}
