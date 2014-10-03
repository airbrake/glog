# Glog

This fork of https://github.com/golang/glog provides all of glog's functionality
and adds the ability to send errors/logs to [Airbrake.io](https://airbrake.io).

## Logging

Please refer to the [glog](https://github.com/golang/glog) code & docs.

## Sending errors to Airbrake.io

A basic example of how to configure glog to send logged errors to Airbrake.io:

```go
package main

import (
  "time"

  "gopkg.in/airbrake/glog.v1"
  "gopkg.in/airbrake/gobrake.v1"
)

var projectID int64 = 123
var apiKey string = "YOUR_API_KEY"

func main() {
  airbrake := gobrake.NewNotifier(projectID, apiKey)
  airbrake.SetContext("environment", "production")
  glog.Gobrake = airbrake

  glog.Error("Gorilla Overflow")

  // Errors are sent asynchronously, allow time for them to send before we exit
  // this example.
  time.Sleep(time.Second * 2)
}
```

## Configure severity

The default is to send only error logs to Airbrake.io. You can change the
severity threshold to also send lower severity logs too, such as warnings:

```go
glog.GobrakeSeverity = glog.WarningLog
```
