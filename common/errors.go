package common

import (
	log "github.com/sirupsen/logrus"
)

func Check(e error) {
	if e != nil {
		log.Panic(e)
	}
}
