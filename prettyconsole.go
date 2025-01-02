package prettyconsole

import (
	"fmt"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorGreen  = "\033[92m"
	ColorBlue   = "\033[34m"
)

var defaultLogColors = LogColors{
	Info:    "\033[34m",
	Success: "\033[92m",
	Warning: "\033[33m",
	Error:   "\033[31m",
	Reset:   "\033[0m",
}

type LogColors struct {
	Reset   string
	Error   string
	Warning string
	Info    string
	Success string
}

type Icons struct {
	Error   string
	Warning string
	Info    string
	Success string
}

var icons = Icons{
	Error:   "❌",
	Warning: "⚠️",
	Info:    "📝",
	Success: "✅",
}

type LoggerOptions struct {
	ShowIcons bool
}

type Logger struct {
	colors  LogColors
	options map[string]bool
}

func NewLogger(options map[string]bool) *Logger {
	loggerOptions := map[string]bool{
		"ShowIcons": true,
	}

	for key, value := range options {
		loggerOptions[key] = value
	}

	return &Logger{
		colors:  defaultLogColors,
		options: loggerOptions,
	}
}

func (logger *Logger) log(logType string, message string) {
	color := logger.getLogTypeColor(logType)
	logIcon := ""

	if logger.options["ShowIcons"] {
		logIcon = "" + getLogIcon(logType) + "  "
	}

	fmt.Printf("%s%s[%s] %s%s\n", color, logIcon, logType, message, ColorReset)
}

func getLogIcon(logType string) string {
	switch logType {
	case "ERROR":
		return icons.Error
	case "WARNING":
		return icons.Warning
	case "INFO":
		return icons.Info
	case "SUCCESS":
		return icons.Success
	default:
		return ""
	}
}

func (logger *Logger) getLogTypeColor(logType string) string {
	switch logType {
	case "ERROR":
		return logger.colors.Error
	case "WARNING":
		return logger.colors.Warning
	case "INFO":
		return logger.colors.Info
	case "SUCCESS":
		return logger.colors.Success
	default:
		return logger.colors.Reset
	}
}

func (logger *Logger) Success(message string) {
	logger.log("SUCCESS", message)
}

func (logger *Logger) Info(message string) {
	logger.log("INFO", message)
}

func (logger *Logger) Warning(message string) {
	logger.log("WARNING", message)
}

func (logger *Logger) Error(message string) {
	logger.log("ERROR", message)
}
