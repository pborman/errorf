package errorf

import (
	"bytes"
	"testing"
)

// Simulate suggested usage
var errf = Errf
var warnf = Warnf

func TestWarnf(t *testing.T) {
	var w bytes.Buffer
	Stderr = &w
	Warnf("message")
	if got, want := w.String(), "message\n"; got != want {
		t.Errorf("Warnf(%q) got %q, want %q", "message", got, want)
	}
	w.Reset()
	Warnf("message\n")
	if got, want := w.String(), "message\n"; got != want {
		t.Errorf("Warnf(%q) got %q, want %q", "message\n", got, want)
	}

	w.Reset()
	Warnf("message %s", "one")
	if got, want := w.String(), "message one\n"; got != want {
		t.Errorf("Warnf(%q, %q) got %q, want %q", "message", "one", got, want)
	}

	w.Reset()
	Warnf("message %s\n", "one")
	if got, want := w.String(), "message one\n"; got != want {
		t.Errorf("Warnf(%q, %q) got %q, want %q", "message\n", "one", got, want)
	}

	w.Reset()
	Warnf("", "one")
	if got, want := w.String(), ""; got != want {
		t.Errorf("Warnf(%q) got %q, want %q", "", got, want)
	}
}

func TestDebugf(t *testing.T) {
	var w bytes.Buffer
	Stderr = &w

	Debug = false
	Debugf("message")
	if got, want := w.String(), ""; got != want {
		t.Errorf("Debugf-false got %q, want %q", got, want)
	}

	w.Reset()
	Debug = true
	Debugf("message")
	if got, want := w.String(), "message\n"; got != want {
		t.Errorf("Debugf-true got %q, want %q", got, want)
	}
}
