package main

import (
	"fmt"
	"os"
)

func getPort() string {
	port := ":8080" // Default port
	if p := os.Getenv("APP_PORT_NUMBER"); p != "" {
		port = ":" + p
	}
	return port
}

func getAppName() (string, error) {
	var appName string
	if name := os.Getenv("APP_NAME"); name != "" {
		appName = name
	} else {
		return "", fmt.Errorf("APP_NAME environment variable is not set")
	}
	return appName, nil
}
