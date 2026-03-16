package main

import (
	"fmt"
	"time"

	"github.com/jakecoffman/cron"
)

func main() {

	fmt.Println(time.Time{})
	c := cron.New()
	c.AddFunc("0 * * * *", func() { fmt.Println("Every second") })
	// c.AddFunc("30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	// c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println(".. in the range 3-6am, 8-11pm") })
	// c.AddFunc("CRON_TZ=Asia/Tokyo 30 04 * * *", func() { fmt.Println("Runs at 04:30 Tokyo time every day") })
	// c.AddFunc("@hourly", func() { fmt.Println("Every hour, starting an hour from now") })
	// c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
	c.Start()
	// c.AddFunc("@daily", func() { fmt.Println("Every day") })
	c.Stop() // Stop the scheduler (does not stop any jobs already running).
	type ParseOption int

	const (
		Second         ParseOption = 1 << iota // Seconds field, default 0
		SecondOptional                         // Optional seconds field, default 0
		Minute                                 // Minutes field, default 0
		Hour                                   // Hours field, default 0
		Dom                                    // Day of month field, default *
		Month                                  // Month field, default *
		Dow                                    // Day of week field, default *
		DowOptional                            // Optional day of week field, default *
		Descriptor                             // Allow descriptors such as @monthly, @weekly, etc.
	)
	fmt.Println(Second)
	fmt.Println(SecondOptional)
	fmt.Println(Minute)
	fmt.Println(Hour)
	var places = []ParseOption{
		Second,
		Minute,
		Hour,
		Dom,
		Month,
		Dow,
	}

	fmt.Println(places)

}
