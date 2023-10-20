package utils

import (
	"log"
	"time"
)

func LogRuntime(begin time.Time) {
	log.Printf("run time is %v nsec", time.Since(begin).Nanoseconds())
}