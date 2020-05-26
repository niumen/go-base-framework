package controller

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func TimeCycle() {
	t := time.NewTicker(time.Duration(time.Second) * 60).C
	for {
		<-t
		log.Debug("time ticker test", time.Now().Format("2006-01-02 15:04:05"))
	}
}
