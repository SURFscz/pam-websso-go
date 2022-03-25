package main

import (
	"fmt"
	"log/syslog"
)


func pamLog(format string, args ...interface{}) {
	l, err := syslog.New(syslog.LOG_AUTH|syslog.LOG_WARNING, "pam-authy")
	if err != nil {
		return
	}
	l.Warning(fmt.Sprintf(format, args...))
}

func main() {

}
