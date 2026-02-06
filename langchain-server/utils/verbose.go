package utils

import (
	"fmt"
	"langchain-mcp-api/env"
)

func IsVerbose() bool {
	return env.IsVerbose
}

func VerbosePrintf(format string, args ...interface{}) {
	if IsVerbose() {
		fmt.Printf(format, args...)
	}
}

func VerbosePrintln(args ...interface{}) {
	if env.IsVerbose {
		fmt.Println(args...)
	}
}
