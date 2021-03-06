// Package errorf provides Go equivalents the functions in <err.h>.
package errorf

import (
	"fmt"
	"io"
	"os"
)

var (
	Stderr = io.Writer(os.Stderr) // Location of standard error
	Debug  = false                // Debugf control
)

// Debugf formats its arguments and writes them to Stderr if Debug is true.  If
// format does not end in a newline a newline is appeneded.  If Debug is false,
// nothing is displayed.
func Debugf(format string, v ...interface{}) {
	if Debug {
		Warnf(format, v...)
	}
}

// Errf formats its arguments and writes them to Stderr and then calls
// os.Exit(1).  If format does not end in a newline a newline is appeneded.
func Errf(format string, v ...interface{}) {
	Warnf(format, v...)
	os.Exit(1)
}

// Warnf formats its arguments and writes them to Stderr.  If format does not
// end in a newline a newline is appeneded.
func Warnf(format string, v ...interface{}) {
	if format == "" {
		return
	}
	if format[len(format)-1] != '\n' {
		format += "\n"
	}
	fmt.Fprintf(Stderr, format, v...)
}
