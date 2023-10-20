package utils

import (
	"log"
	"runtime"
)

func MemoryAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Memory: %d KB\n", m.Alloc/1024)
}
