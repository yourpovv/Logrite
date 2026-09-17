package main

import (
	logger "github.com/yourpovv/logrite"
)

func main() {
	logger.SetConfig(logger.Config{
		ShowIcons:    true,
		UppercaseTag: true,
		UseColors:    true,
	})

	logger.Info("This is an info log")
	logger.Warn("This is a warning log")
	logger.Error("This is an error log")
	logger.Success("This is a success log")
	logger.Debug("This is a debug log\n")

	logger.WebStart(8080)
	logger.Cmd("YourPOV", "Clear")
	logger.Login("admin", "127.0.0.1\n")

	logger.Custom("🥂", "custom", "This is a custom log with a custom emoji", logger.Green, logger.BgWhite)

	logger.Log("manual", " This is a manually made log\n", logger.BgBlack, logger.BrightWhite)
}
