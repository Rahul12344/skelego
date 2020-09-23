package skelego

import (
	"fmt"
	"log"
	"sync"
)

//Logging Logging functions necessary for application,
type Logging interface {
	LogEvent(string, ...interface{})
	LogError(string, ...interface{})
	LogFatal(string, ...interface{})
}

type logger struct {
	lock sync.RWMutex
}

//NewLogger Creates a new logger - should have exactly one logger per application.
func NewLogger() Logging {
	return &logger{}
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
