package utils

import (
	"os"

	"github.com/gcp-kit/line-bot-boilerplate-go-example/pkg/constant"
)

var (
	projectID      = getProjectID()
	serviceName    = os.Getenv(constant.EnvKeyServiceName)
	serviceVersion = os.Getenv(constant.EnvKeyServiceVersion)
)

func getProjectID() string {
	id, ok := os.LookupEnv(constant.EnvKeyProjectID)
	if ok {
		return id
	}
	id, ok = os.LookupEnv(constant.EnvKeyGoogleCloudProject)
	if ok {
		return id
	}
	return ""
}

// GetProjectID - get the project id
func GetProjectID() string {
	id := projectID
	if id == "" {
		return "localProject"
	}
	return id
}

// GetServiceName - get the service name
func GetServiceName() string {
	if serviceName == "" {
		return "localService"
	}
	return serviceName
}

// GetServiceVersion - get the version
func GetServiceVersion() string {
	if serviceVersion == "" {
		return "1.0"
	}
	return serviceVersion
}
