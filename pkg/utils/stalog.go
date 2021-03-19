package utils

import (
	"os"

	"github.com/gcp-kit/stalog"
)

// NewStalogConfig ...
func NewStalogConfig(severities ...stalog.Severity) *stalog.Config {
	severity := stalog.SeverityDefault
	if len(severities) > 0 {
		severity = severities[0]
	}

	config := stalog.NewConfig(GetProjectID())
	config.RequestLogOut = os.Stderr               // request log to stderr
	config.ContextLogOut = os.Stdout               // context log to stdout
	config.Severity = severity                     // only over variable `severity` logs are logged
	config.AdditionalData = stalog.AdditionalData{ // set additional fields for all logs
		"service": GetServiceName(),
		"version": GetServiceVersion(),
	}
	return config
}
