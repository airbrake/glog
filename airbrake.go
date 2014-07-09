package glog

import (
	"fmt"

	"gopkg.in/airbrake/gobrake.v1"
)

var Gobrake *gobrake.Notifier

func notifyAirbrake(s severity, format string, args ...interface{}) {
	if Gobrake == nil {
		return
	}
	if s < errorLog {
		return
	}

	var msg string
	if format != "" {
		msg = fmt.Sprintf(format, args...)
	} else {
		msg = fmt.Sprint(args...)
	}

	foundErr := false
	for _, arg := range args {
		err, ok := arg.(error)
		if !ok {
			continue
		}
		foundErr = true

		notice := Gobrake.Notice(err, nil, 5)
		notice.Env["glog_message"] = msg
		go Gobrake.SendNotice(notice)
	}

	if !foundErr {
		notice := Gobrake.Notice(msg, nil, 5)
		go Gobrake.SendNotice(notice)
	}
}
