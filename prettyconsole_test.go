package prettyconsole

import "testing"

func TestLogger(t *testing.T) {
	logger := NewLogger()

	logger.Success("This is a success message.")
	logger.Info("This is an info message.")
	logger.Warning("This is a warning message.")
	logger.Error("This is an error message.")
}
