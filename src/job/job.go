package job

import (
	"main/config"
	"main/pkg/data"
	"strconv"
	"time"
)

func TimedJob(ticker *time.Ticker, done chan bool) {
	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case <-ticker.C:
			data.DumpDataToFile()
		}
	}
}
func TimedJobStart() {
	duration, _ := strconv.Atoi(config.DUMP_DURATION)
	ticker := time.NewTicker(time.Duration(duration) * time.Second)
	done := make(chan bool)
	go TimedJob(ticker, done)
}
