package util

import (
	"runtime"

	log "github.com/sirupsen/logrus"
)

// Catch recovers from the error
func Catch(err *error) {
	if r := recover(); r != nil {
		panicError := r.(error)
		*err = panicError
	}
}

// Check panics and log the error if any
func Check(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Errorf("file (%v) line: %v got error: %v", file, line, err.Error())
		panic(err)
	}
}
