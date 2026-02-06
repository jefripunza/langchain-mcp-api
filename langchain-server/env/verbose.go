package env

import (
	"os"
	"strings"
)

var IsVerbose bool
var Verbose string

func init() {
	Verbose := strings.ToLower(os.Getenv("VERBOSE"))
	IsVerbose = Verbose == "true" || Verbose == "1"
	// isVerbose = true // TODO: remove this
}
