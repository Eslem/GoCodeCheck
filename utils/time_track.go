package utils

import (
	"log"
	"time"
)

// TimeTrack calculate the time and print it
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s  \n", name, elapsed)
}
