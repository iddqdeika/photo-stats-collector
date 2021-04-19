package processor

import "photo-stats-collector/definitions"

func New(id definitions.ImageDownloader, sc definitions.StatCalculator) definitions.Processor {
	return &processor{
		id: id,
		sc: sc,
	}
}

type processor struct {
	id definitions.ImageDownloader
	sc definitions.StatCalculator
}

func (p *processor) Process(url definitions.ImageUrl) definitions.ImageStats {
	e := func(err error) definitions.ImageStats {
		return definitions.ImageStats{Url: string(url), Err: err}
	}
	data, err := p.id.Download(url)
	if err != nil {
		return e(err)
	}
	stats, err := p.sc.GetStats(data)
	if err != nil {
		return e(err)
	}
	stats.Url = string(url)
	return *stats
}
