package util

import "fmt"

func Sprint(format interface{}) string {
	return fmt.Sprint(format)
}

func Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}
