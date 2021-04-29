package skelego

import (
	"fmt"
	"log"
)

var (
	globalLogger *logger
)

type logger struct {
	name string
}

//NewLogger Creates a new logger - should have exactly one logger per application.
func NewLogger() *logger {
	return &logger{
		name: "Logger",
	}
}

//Logs event.
func (logs *logger) LogEvent(val string, args ...interface{}) {
	log.Printf(val, args...)
}

//Logs non-fatal error.
func (logs *logger) LogError(val string, args ...interface{}) {
	log.Printf(fmt.Sprintf("Error: %s", val), args...)
}

//Logs fatal event and terminates.
func (logs *logger) LogFatal(val string, args ...interface{}) {
	log.Fatalf(val, args...)
}

//New Logger
func Logger() *logger {
	if globalLogger == nil {
		globalLogger = NewLogger()
	}
	return globalLogger
}
