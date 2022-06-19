package singleton

import (
	"log"
	"os"
	"sync"
)

// Logger is a logger type and an actual usecase for this pattern.
type Logger struct {
	Info  *log.Logger
	Error *log.Logger
}

var (
	once     *sync.Once
	instance *Logger
)

func init() {
	once.Do(func() {
		instance = &Logger{
			Info:  log.New(os.Stdout, "info", log.LstdFlags),
			Error: log.New(os.Stderr, "error", log.LstdFlags),
		}
	})
}

func GetAppLogger() *Logger {
	return instance
}
