package main

import (
	"openingjobs/pkg/config"
	"openingjobs/pkg/context"
)

func main() {
	var err error
	logger := config.GetLogger()

	// Initialize config
	err = config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	defer config.GetDB().Close()

	// Initialize handler contexts
	err = context.InitializeHandlerContexts()
	if err != nil {
		logger.Errorf("context initialization error: %v", err)
		return
	}
}
