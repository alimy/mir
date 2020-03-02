package version

import "github.com/coreos/go-semver/semver"

var (
	// GitHash Value will be set during build
	GitHash = "Not provided"
	// BuildTime Value will be set during build
	BuildTime = "Not provided"
)

// MircVer version of Mirc
var MircVer = semver.Version{
	Major: 0,
	Minor: 1,
	Patch: 0,
}
