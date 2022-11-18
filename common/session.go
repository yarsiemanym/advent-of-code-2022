package common

import "os"

var Session string

func InitSession() {
	session := os.Getenv("AOC_SESSION_TOKEN")
	Session = session
}
