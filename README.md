# PrettyConsole

PrettyConsole is a Go package for logging with colorized output for different log levels (success, info, warning, error). It provides a simple and clean way to display logs in the terminal, enhancing readability.

## Features
- **Colorized Logs**: Logs are displayed with distinct colors for better readability.
- **Simple API**: Easy-to-use API for logging at different levels.
- **Supports Multiple Log Levels**: Supports `SUCCESS`, `INFO`, `WARNING`, and `ERROR` log levels.

## Installation

To use PrettyConsole in your Go project, run:

```bash
go get github.com/unalsemih/prettyconsole
```

## Usage

```go
package main

import (
	"github.com/unalsemih/prettyconsole"
)

func main() {
	logger := prettyconsole.NewLogger(map[string]bool{})

	logger.Success("This is a success message.")
	logger.Info("This is an info message.")
	logger.Warning("This is a warning message.")
	logger.Error("This is an error message.")
}
```

## Licence

The MIT License (MIT) - see [`LICENSE`](https://github.com/unalsemih/prettyconsole/blob/main/LICENSE) for more details
