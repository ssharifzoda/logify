package logify

import (
	"runtime"
	"strings"
)

func GetLogLevel(level LogLevel) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case FatalLevel:
		return "FATAL"
	case ErrorLevel:
		return "ERROR"
	case InfoLevel:
		return "INFO"
	case WarningLevel:
		return "WARNING"
	default:
		return "UNKNOWN"
	}
}

// GetCaller возвращает информацию о вызывающей функции.
func GetCaller() (string, int) {
	// Получаем информацию о стеке вызовов
	var pc [16]uintptr
	n := runtime.Callers(4, pc[:])
	frames := runtime.CallersFrames(pc[:n])

	// Проходим по стеку вызовов и находим первую функцию, не являющуюся функцией логирования
	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.Function, "logger.") {
			return frame.File, frame.Line
		}
		if !more {
			break
		}
	}
	// Если не удалось найти вызывающую функцию, возвращаем пустые значения
	return "", 0
}
