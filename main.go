package main

import (
	"github.com/iddqdeika/rrr"
	"photo-stats-collector/root"
)

func main() {
	rrr.BasicEntry(root.New())
}
