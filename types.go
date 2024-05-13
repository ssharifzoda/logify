package logify

import (
	"os"
	"sync"
	"time"
)

type LogLevel int

type Logger struct {
	file   *os.File
	lock   sync.Mutex
	level  LogLevel
	format func(level LogLevel, time time.Time, file string, line int, message string) string
}

type LogEntry struct {
	Time     time.Time `json:"time"`
	Level    LogLevel  `json:"level"`
	Filename string    `json:"filename"`
	Line     int       `json:"line"`
	Message  string    `json:"message"`
}
