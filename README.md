# errorf

A Go version of C's <err.h>

This pacakge provides 3 functions, Debugf, Errf, and Warnf.  They all write formatted output to standard error.  The differences are:

* Debugf - is a no-op if errf.Debug is false
* Errf - calls os.Exit(1) instead of returning
* Warnf - nothing special, just displays the output

To reduce the amount of typing, it is suggested that these values be aliased.

import "github.com/pborman/errorf"

var errf = errorf.Errf
var warnf = errorf.Warnf
