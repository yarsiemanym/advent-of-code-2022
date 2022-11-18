package common

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogging() {
	envVar := os.Getenv("AOC_LOG_LEVEL")

	if envVar != "" {
		logLevel, err := log.ParseLevel(os.Getenv("AOC_LOG_LEVEL"))

		if err != nil {
			log.SetLevel(log.WarnLevel)
			log.Warnf("\"%v\" is not a valid log level. Falling back to \"warn\".", envVar)
		} else {
			log.SetLevel(logLevel)
		}
	} else {
		log.SetLevel(log.WarnLevel)
	}
}
