package glog

import (
	"fmt"
	"net/http"

	"gopkg.in/airbrake/gobrake.v1"
)

// Gobrake is an instance of Airbrake Go Notifier that is used to send
// logs to Airbrake.
var Gobrake *gobrake.Notifier

// Minimum log severity that will be sent to Airbrake.
var GobrakeSeverity = ErrorLog

type requester interface {
	Request() *http.Request
}

func notifyAirbrake(s severity, format string, args ...interface{}) {
	if Gobrake == nil {
		return
	}
	if s < GobrakeSeverity {
		return
	}

	var msg string
	if format != "" {
		msg = fmt.Sprintf(format, args...)
	} else {
		msg = fmt.Sprint(args...)
	}

	var req *http.Request
	for _, arg := range args {
		if v, ok := arg.(requester); ok {
			req = v.Request()
			break
		}
	}

	foundErr := false
	for _, arg := range args {
		err, ok := arg.(error)
		if !ok {
			continue
		}
		foundErr = true

		notice := Gobrake.Notice(err, req, 4)
		notice.Env["glog_message"] = msg
		go Gobrake.SendNotice(notice)
	}

	if !foundErr {
		notice := Gobrake.Notice(msg, req, 4)
		go Gobrake.SendNotice(notice)
	}
}
