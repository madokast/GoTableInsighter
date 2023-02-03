package main

import (
	"tableinsight/internal/logger"
)

func main() {
	logger.Debug("Hello", "123")
	logger.Info("Hello", 123)
	logger.Warn("Hello", 3.14)
	logger.Error("Hello", false)
}
