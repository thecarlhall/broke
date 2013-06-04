// Package logger contains functionality related to writing to a log.
package logger

import (
	"fmt"
	"log"
)

const (
	// Log levels
	TRACE = 1
	DEBUG = 2
	INFO  = 3
	WARN  = 4
	ERROR = 5

	LEVEL = INFO
)

// Tracef logs a formatted message when the logging level is at least TRACE.
// Arguments are handled in the manner of log.Printf.
func Tracef(format string, v ...interface{}) {
	if LEVEL <= TRACE {
		Printf("[TRACE]", format, v...)
	}
}

// Trace logs a formatted message when the logging level is at least TRACE.
// Arguments are handled in the manner of log.Printf.
func Trace(v ...interface{}) {
	if LEVEL <= TRACE {
		Print("[TRACE]", v...)
	}
}

// Debugf logs a formatted message when the logging level is at least DEBUG.
// Arguments are handled in the manner of log.Printf.
func Debugf(format string, v ...interface{}) {
	if LEVEL <= DEBUG {
		Printf("[DEBUG]", format, v...)
	}
}

// Debug logs a formatted message when the logging level is at least DEBUG.
// Arguments are handled in the manner of log.Printf.
func Debug(v ...interface{}) {
	if LEVEL <= DEBUG {
		Print("[DEBUG]", v...)
	}
}

// Infof logs a formatted message when the logging level is at least INFO.
// Arguments are handled in the manner of log.Printf.
func Infof(format string, v ...interface{}) {
	if LEVEL <= INFO {
		Printf("[INFO]", format, v...)
	}
}

// Infof logs when the logging level is at least INFO.
// Arguments are handled in the manner of log.Print.
func Info(v ...interface{}) {
	if LEVEL <= INFO {
		Print("[INFO]", v...)
	}
}

// Warnf logs a formatted message when the logging level is at least WARN.
// Arguments are handled in the manner of log.Printf.
func Warnf(format string, v ...interface{}) {
	if LEVEL <= WARN {
		Printf("[WARN]", format, v...)
	}
}

// Warn logs a formatted message when the logging level is at least WARN.
// Arguments are handled in the manner of log.Printf.
func Warn(v ...interface{}) {
	if LEVEL <= WARN {
		Print("[WARN]", v...)
	}
}

// Errorf logs a formatted message when the logging level is at least ERROR.
// Arguments are handled in the manner of log.Printf.
func Errorf(format string, v ...interface{}) {
	if LEVEL <= ERROR {
		Printf("[ERROR]", format, v...)
	}
}

// Error logs when the logging level is at least ERROR.
// Arguments are handled in the manner of log.Print.
func Error(v ...interface{}) {
	if LEVEL <= ERROR {
		Print("[ERROR]", v...)
	}
}

// Fatal is a delegation to Go's log.Fatalf. This is just for consistency.
func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

// Fatal is a delegation to Go's log.Fatal. This is just for consistency.
func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

// Print calls the Go log to write output.
// Arguments are handled in the manner of log.Printf.
func Printf(prefix string, message string, v ...interface{}) {
	log.Printf("%s %s\n", prefix, fmt.Sprintf(message, v...))
}

// Print calls the Go log to write output.
// Arguments are handled in the manner of log.Print.
func Print(prefix string, v ...interface{}) {
	log.Printf("%s %s\n", prefix, fmt.Sprint(v...))
}
