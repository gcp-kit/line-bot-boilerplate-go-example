package constant

import (
	"os"
)

const (
	EnvKeyProjectID          = "GCP_PROJECT"          // for Google Cloud Functions
	EnvKeyGoogleCloudProject = "GOOGLE_CLOUD_PROJECT" // for Appengine
	EnvKeyServiceName        = "GAE_SERVICE"
	EnvKeyServiceVersion     = "GAE_VERSION"
)

// CommitHash - get the commit hash(short)
var CommitHash = func() string {
	hash, ok := os.LookupEnv("SHORT_SHA")
	if !ok {
		return "local"
	}
	return hash
}()
