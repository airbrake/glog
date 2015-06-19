// github.com/golang/glog fork with Airbrake integration.
//
// Copyright 2013 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package glog

import (
	"fmt"

	golog "github.com/golang/glog"
)

// severity identifies the sort of log: info, warning etc. It also implements
// the flag.Value interface. The -stderrthreshold flag is of type severity and
// should be modified only through the flag.Value interface. The values match
// the corresponding constants in C++.
type severity int32 // sync/atomic int32

// These constants identify the log levels in order of increasing severity.
// A message written to a high-severity log file is also written to each
// lower-severity log file.
const (
	InfoLog severity = iota
	WarningLog
	ErrorLog
	FatalLog
	numSeverity = 4
)

// Stats tracks the number of lines of output and number of bytes
// per severity level. Values must be read with atomic.LoadInt64.
var Stats = golog.Stats

// SetMaxSize sets maximum size of a log file in bytes.
func SetMaxSize(maxSize uint64) {
	golog.MaxSize = maxSize
}

// CopyStandardLogTo arranges for messages written to the Go "log" package's
// default logs to also appear in the Google logs for the named and lower
// severities.  Subsequent changes to the standard log's default output location
// or format may break this behavior.
//
// Valid names are "INFO", "WARNING", "ERROR", and "FATAL".  If the name is not
// recognized, CopyStandardLogTo panics.
func CopyStandardLogTo(name string) {
	golog.CopyStandardLogTo(name)
}

// Flush flushes all pending log I/O.
func Flush() {
	golog.Flush()
}

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Info(args ...interface{}) {
	msg := fmt.Sprint(args...)
	golog.InfoDepth(1, msg)
	notifyAirbrake(InfoLog, msg)
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Infoln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	golog.InfoDepth(1, msg)
	notifyAirbrake(InfoLog, msg)
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	golog.InfoDepth(1, msg)
	notifyAirbrake(InfoLog, msg)
}

// Warning logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Warning(args ...interface{}) {
	msg := fmt.Sprint(args...)
	golog.WarningDepth(1, msg)
	notifyAirbrake(WarningLog, msg)
}

// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Warningln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	golog.WarningDepth(1, msg)
	notifyAirbrake(WarningLog, msg)
}

// Warningf logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Warningf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	golog.WarningDepth(1, msg)
	notifyAirbrake(WarningLog, msg)
}

// Error logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Error(args ...interface{}) {
	msg := fmt.Sprint(args...)
	golog.ErrorDepth(1, msg)
	notifyAirbrake(ErrorLog, msg)
}

// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Errorln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	golog.ErrorDepth(1, msg)
	notifyAirbrake(ErrorLog, msg)
}

// Errorf logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	golog.ErrorDepth(1, msg)
	notifyAirbrake(ErrorLog, msg)
}

// Fatal logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Fatal(args ...interface{}) {
	msg := fmt.Sprint(args...)
	golog.FatalDepth(1, msg)
	notifyAirbrake(FatalLog, msg)
}

// Fatalln logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Fatalln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	golog.FatalDepth(1, msg)
	notifyAirbrake(FatalLog, msg)
}

// Fatalf logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Fatalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	golog.FatalDepth(1, msg)
	notifyAirbrake(FatalLog, msg)
}

// Exit logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Exit(args ...interface{}) {
	msg := fmt.Sprint(args...)
	golog.ExitDepth(1, msg)
	notifyAirbrake(ErrorLog, msg)
}

// Exitln logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
func Exitln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	golog.ExitDepth(1, msg)
	notifyAirbrake(ErrorLog, msg)
}

// Exitf logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Exitf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	golog.ExitDepth(1, msg)
	notifyAirbrake(ErrorLog, msg)
}

// Verbose is a boolean type that implements Infof (like Printf) etc.
// See the documentation of V for more information.
type Verbose bool

// V reports whether verbosity at the call site is at least the requested level.
// The returned value is a boolean of type Verbose, which implements Info, Infoln
// and Infof. These methods will write to the Info log if called.
// Thus, one may write either
//	if glog.V(2) { glog.Info("log this") }
// or
//	glog.V(2).Info("log this")
// The second form is shorter but the first is cheaper if logging is off because it does
// not evaluate its arguments.
//
// Whether an individual call to V generates a log record depends on the setting of
// the -v and --vmodule flags; both are off by default. If the level in the call to
// V is at least the value of -v, or of -vmodule for the source file containing the
// call, the V call will log.
func V(level golog.Level) Verbose {
	return Verbose(golog.V(level))
}

// Info is equivalent to the global Info function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) Info(args ...interface{}) {
	if v {
		Info(args...)
	}
}

// Infoln is equivalent to the global Infoln function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) Infoln(args ...interface{}) {
	if v {
		Infoln(args...)
	}
}

// Infof is equivalent to the global Infof function, guarded by the value of v.
// See the documentation of V for usage.
func (v Verbose) Infof(format string, args ...interface{}) {
	if v {
		Infof(format, args...)
	}
}
