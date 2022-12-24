package main

import (
	"fmt"
	"os"
)

type YouTrack struct {
	URL   string
	Token string
}

func loadConfig() (*YouTrack, error) {
	config := &YouTrack{}
	for _, env := range []string{"YOUTRACK_URL", "YOUTRACK_TOKEN"} {
		value := os.Getenv(env)
		if value == "" {
			return config, fmt.Errorf("empty value for %s", env)
		}
		switch env {
		case "YOUTRACK_URL":
			config.URL = env
		case "YOUTRACK_TOKEN":
			config.Token = env
		}
	}

	return config, nil
}
