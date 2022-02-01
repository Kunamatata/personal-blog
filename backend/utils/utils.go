package utils

import (
	"log"
	"time"
)

func Duration(msg string, start time.Time) {
	log.Printf("%s took %s", msg, time.Since(start))
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}
