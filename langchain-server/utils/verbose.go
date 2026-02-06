package utils

import (
	"fmt"
	"os"
	"strings"
)

var isVerbose bool

func init() {
	verbose := strings.ToLower(os.Getenv("VERBOSE"))
	isVerbose = verbose == "true" || verbose == "1"
	isVerbose = true // TODO: remove this
}

func IsVerbose() bool {
	return isVerbose
}

func VerbosePrintf(format string, args ...interface{}) {
	if isVerbose {
		fmt.Printf(format, args...)
	}
}

func VerbosePrintln(args ...interface{}) {
	if isVerbose {
		fmt.Println(args...)
	}
}
