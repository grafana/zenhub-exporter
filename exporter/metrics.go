package exporter

import (
	"github.com/grafana/zenhub-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
)

// RateLimitExceededStatus is the status response from zenhub when the rate limit is exceeded.
const RateLimitExceededStatus = "403 rate limit exceeded"

// Exporter is used to store Metrics data and embeds the config struct.
type Exporter struct {
	APIMetrics map[string]*prometheus.Desc
	config.Config
}
