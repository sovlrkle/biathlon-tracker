package main

import (
	"biathlon-tracker/config"
	"biathlon-tracker/internal/app"
	"fmt"
	"os"
)

func main() {
	eventsPath := os.Getenv("EVENTS_PATH")
	if eventsPath == "" {
		fmt.Println("EVENTS_PATH environment variable is not set")
		return
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		fmt.Println("CONFIG_PATH environment variable is not set")
		return
	}

	cfg, err := config.New(configPath)
	if err != nil {
		fmt.Printf("Failed to parse config: %v\n", err)
		return
	}

	if err = app.Run(cfg, eventsPath); err != nil {
		fmt.Printf("Application error: %v\n", err)
	}

}
