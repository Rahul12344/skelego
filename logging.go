package skelego

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var (
	globalLogger *logger
)

type logger struct {
	logfilename string
	logfile     *os.File
	log         *log.Logger
	lock        sync.RWMutex
}

//NewLogger Creates a new logger - should have exactly one logger per application.
func NewLogger() *logger {
	file, _ := os.Create("logfile")
	return &logger{
		logfilename: "logfile",
		logfile:     file,
		log:         log.New(file, "recipes", log.Lshortfile),
	}
}

//Logs event.
func (logs *logger) LogEvent(val string, args ...interface{}) {
	logs.lock.RLock()
	logs.log.Printf(val, args...)
	logs.lock.RUnlock()
}

//Logs non-fatal error.
func (logs *logger) LogError(val string, args ...interface{}) {
	logs.lock.RLock()
	logs.log.Printf(fmt.Sprintf("Error: %s", val), args...)
	logs.lock.RUnlock()
}

//Logs fatal event and terminates.
func (logs *logger) LogFatal(val string, args ...interface{}) {
	logs.lock.RLock()
	logs.log.Fatalf(val, args...)
	logs.lock.RUnlock()
}

// Perform the actual act of rotating and reopening file.
func (logs *logger) Rotate() (err error) {
	logs.lock.Lock()
	defer logs.lock.Unlock()

	// Close existing file if open
	if logs.logfile != nil {
		err = logs.logfile.Close()
		logs.logfile = nil
		if err != nil {
			return
		}
	}
	// Rename dest file if it already exists
	_, err = os.Stat(logs.logfilename)
	if err == nil {
		err = os.Rename(logs.logfilename, logs.logfilename+"."+time.Now().Format(time.RFC3339))
		if err != nil {
			return
		}
	}

	// Create a file.
	logs.logfile, err = os.Create(logs.logfilename)
	return
}

//New Logger
func Logger() *logger {
	if globalLogger == nil {
		globalLogger = NewLogger()
	}
	return globalLogger
}
