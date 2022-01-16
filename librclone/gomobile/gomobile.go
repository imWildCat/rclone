// Package gomobile exports shims for gomobile use
package gomobile

import (
	_ "github.com/rclone/rclone/backend/all" // import all backends
	"github.com/rclone/rclone/fs/config"
	"github.com/rclone/rclone/lib/oauthutil"
	_ "github.com/rclone/rclone/lib/plugin" // import plugins
	"github.com/rclone/rclone/librclone/librclone"
)

// RcloneInitialize initializes rclone as a library
func RcloneInitialize() {
	librclone.Initialize()
}

func SetConfigPath(path string) {
	config.SetConfigPath(path)
}

// define call back interface
type StringCallback interface {
	Callback(str string)
}

func SetAuthCallback(callback StringCallback) {
	oauthutil.SetAuthCallbackFunc(callback.Callback)
}

// RcloneFinalize finalizes the library
func RcloneFinalize() {
	librclone.Finalize()
}

// RcloneRPCResult is returned from RcloneRPC
//
//   Output will be returned as a serialized JSON object
//   Status is a HTTP status return (200=OK anything else fail)
type RcloneRPCResult struct {
	Output string
	Status int
}

// RcloneRPC has an interface optimised for gomobile, in particular
// the function signature is valid under gobind rules.
//
// https://pkg.go.dev/golang.org/x/mobile/cmd/gobind#hdr-Type_restrictions
func RcloneRPC(method string, input string) (result *RcloneRPCResult) { //nolint:deadcode
	output, status := librclone.RPC(method, input)
	return &RcloneRPCResult{
		Output: output,
		Status: status,
	}
}
