package logify

const (
	logsDir  = "logs"
	logsFile = "all.log"
)

const (
	// TraceLevel используется для трассировки наименее значимых событий.
	TraceLevel LogLevel = iota
	// DebugLevel используется для отладочных сообщений.
	DebugLevel
	// InfoLevel используется для информационных сообщений.
	InfoLevel
	// WarningLevel используется для предупреждений.
	WarningLevel
	// ErrorLevel используется для сообщений об ошибках.
	ErrorLevel
	// FatalLevel используется для критических ошибок, которые приводят к завершению программы.
	FatalLevel
)
