package prettyconsole

import "fmt"

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorGreen  = "\033[92m"
	ColorBlue   = "\033[34m"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (logger *Logger) Success(message string) {
	fmt.Printf("%s[SUCCESS] %s%s\n", ColorGreen, message, ColorReset)
}

func (logger *Logger) Info(message string) {
	fmt.Printf("%s[INFO] %s%s\n", ColorBlue, message, ColorReset)
}

func (logger *Logger) Warning(message string) {
	fmt.Printf("%s[WARNING] %s%s\n", ColorYellow, message, ColorReset)
}

func (logger *Logger) Error(message string) {
	fmt.Printf("%s[ERROR] %s%s\n", ColorRed, message, ColorReset)
}
