package main

import (
	"fmt"
	"log/syslog"
	//"net/http"
	//"io/ioutil"
	//"encoding/json"
)


func pamLog(format string, args ...interface{}) {
	l, err := syslog.New(syslog.LOG_AUTH|syslog.LOG_WARNING, "pam-authy")
	if err != nil {
		return
	}
	l.Warning(fmt.Sprintf(format, args...))
}

type Req struct {
	nonce string
	pin string
	challenge string
	hot bool
}

type Auth struct {
	uid string
	success	bool
}

func (req *Req) startReq(userID string) (*Req, error) {

	req.nonce = "1234"
	req.pin = "4321"
	req.challenge = "Press enter"
	req.hot = false

	return req, nil
}

func (auth *Auth) getAuth(nonce string) (*Auth, error) {
	auth.uid = "Test"
	auth.success = true

	return auth, nil
}

func main() {

}
