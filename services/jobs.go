package services

import (
	"time"

	"github.com/go-co-op/gocron"
)

func RunCronJobs() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Minute().Do(func() {
		// * running service if server down
		start, end := getStartEnd()

		IfServerShutdown(start, end)
	})

	s.StartBlocking()
}

func getStartEnd() (start int, end int) {
	timeNow := int(time.Now().Unix())

	start = timeNow - 59
	end = timeNow

	return start, end
}
