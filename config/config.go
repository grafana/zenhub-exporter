package config

import (
	"os"

	cfg "github.com/infinityworks/go-common/config"
	"github.com/prometheus/common/log"
)

// Config holds the runtime configuration for the exporter
type Config struct {
	*cfg.BaseConfig
	APIURL      string
	RepoID      string
	WorkspaceID string
	APIToken    string
}

// Init populates the configuration at runtime
func Init() Config {

	listenPort := cfg.GetEnv("LISTEN_PORT", "9171")
	os.Setenv("LISTEN_PORT", listenPort)
	ac := cfg.Init()
	url := cfg.GetEnv("API_URL", "https://api.zenhub.com")
	repoID := os.Getenv("REPO_ID")
	workspaceID := os.Getenv("WORKSPACE_ID")
	apiToken := os.Getenv("API_TOKEN")

	if repoID == "" {
		log.Error("REPO_ID environment variable not set... exiting")
	}

	if workspaceID == "" {
		log.Error("WORKSPACE_ID environment variable not set... exiting")
	}

	if apiToken == "" {
		log.Error("API_TOKEN environment variable not set... exiting")
	}

	appConfig := Config{
		&ac,
		url,
		repoID,
		workspaceID,
		apiToken,
	}

	return appConfig
}
