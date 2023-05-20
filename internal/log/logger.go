package log

import (
	"fmt"
	"log"
)

func Debugf(format string, args ...interface{}) {
	logFormat := fmt.Sprintf("[DEBUG] %s", format)
	log.Printf(logFormat, args...)
}

func Infof(format string, args ...interface{}) {
	logFormat := fmt.Sprintf("[INFO] %s", format)
	log.Printf(logFormat, args...)
}

func Warnf(format string, args ...interface{}) {
	logFormat := fmt.Sprintf("[WARN] %s", format)
	log.Printf(logFormat, args...)
}

func Errorf(format string, args ...interface{}) {
	logFormat := fmt.Sprintf("[ERROR] %s", format)
	log.Printf(logFormat, args...)
}
