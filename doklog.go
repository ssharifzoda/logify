package logify

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// NewLogger создает новый экземпляр logify-a
func NewLogger(file *os.File, level LogLevel, format func(level LogLevel, time time.Time, file string, line int, message string) string) *Logger {
	return &Logger{
		file:   file,
		level:  level,
		format: format,
	}
}

// Init инициализирует логгер с указанными параметрами.
// Пример:
// log := logify.Init("log", "all.log", logify.InfoLevel, logify.DefaultLogFormat)
//
//	defer log.Close()
func Init(logsDir, logsFile string, level LogLevel, format func(level LogLevel, time time.Time, file string, line int, message string) string) *Logger {
	logsPath := filepath.Join(logsDir, logsFile)
	err := os.MkdirAll(logsDir, os.ModePerm) // Создаем директорию, если она не существует
	if err != nil {
		fmt.Println("Ошибка при создании директории логов:", err)
		return nil
	}
	file, err := os.OpenFile(logsPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ошибка при открытии файла логов:", err)
		return nil
	}
	logger := NewLogger(file, level, format)
	if logger == nil {
		fmt.Println("Ошибка при создании логгера")
		return nil
	}
	return logger
}

// Log записывает сообщение в журнал с указанным уровнем
func (l *Logger) Log(level LogLevel, format string, args ...interface{}) {
	if level < l.level {
		return
	}
	l.lock.Lock()
	defer l.lock.Unlock()

	logTime := time.Now()
	file, line := GetCaller() // Получаем информацию о вызывающей функции
	fileName := filepath.Base(file)

	logEntry := &LogEntry{
		Time:     logTime,
		Level:    level,
		Filename: fileName,
		Line:     line,
		Message:  fmt.Sprintf(format, args...),
	}

	logOutput := l.format(logEntry.Level, logEntry.Time, logEntry.Filename, logEntry.Line, logEntry.Message)
	_, err := l.file.WriteString(logOutput + "\n")
	if err != nil {
		fmt.Println("Ошибка при записи в журнал:", err)
	}
}

func (l *Logger) Fatal(args ...interface{}) {
	l.Log(FatalLevel, "%s", args...)
	l.Close()
	os.Exit(1)
}

func (l *Logger) Info(args ...interface{}) {
	l.Log(InfoLevel, "%s", args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.Log(ErrorLevel, "%s", args...)
}

func (l *Logger) Warning(args ...interface{}) {
	l.Log(WarningLevel, "%s", args...)
}

// Close закрывает файл журнала.
func (l *Logger) Close() {
	if l.file != nil {
		l.file.Close()
	}
}

// DefaultLogFormat форматирует запись журнала по умолчанию.
func DefaultLogFormat(level LogLevel, time time.Time, file string, line int, message string) string {
	return fmt.Sprintf("[%s]-[%s] %s:%d - %s", time.Format("2006-01-02 15:04:05"),
		levelString(level), file, line, message)
}

// levelString возвращает строковое представление уровня журнала.
func levelString(level LogLevel) string {
	switch level {
	case TraceLevel:
		return "TRACE"
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarningLevel:
		return "WARNING"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Примеры использования:

// Пример использования Init и Log:
// log := logger.Init("logs", "app.log", logger.InfoLevel, logger.DefaultLogFormat)
// log.Log(logger.InfoLevel, "This is an info message")

// Пример использования Close:
// log.Close()
