package main

import (
	"github.com/robfig/cron/v3"
	"fmt"
	"time"
)

func main() {
	c := cron.New()
	c.AddFunc("@every 0h1m", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
	c.Start()
	time.Sleep(5 * time.Minute)
}

