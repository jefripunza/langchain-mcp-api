package utils

import "time"

func Delay(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func StringPtr(s string) *string {
	return &s
}

func IntPtr(i int) *int {
	return &i
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
