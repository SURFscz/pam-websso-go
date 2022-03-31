package main

import (
	"fmt"
	"log/syslog"
	"net/http"
	"net/url"
	"log"
	//"io/ioutil"
	//"encoding/json"
)


func pamLog(format string, args ...interface{}) {
	log.Printf(fmt.Sprintf(format, args...))

	return

	l, err := syslog.New(syslog.LOG_AUTH|syslog.LOG_WARNING, "pam-authy")
	if err != nil {
		return
	}
	l.Warning(fmt.Sprintf(format, args...))
}

type Req struct {
	Nonce string		`json:"nonce"`
	Pin string		`json:"pin"`
	Challenge string	`json:"challenge"`
	Hot bool		`json:"hot"`
}

func (req *Req) startReq(user string) error {
	pamLog("Start request for: %s", user)

	sso_url := "http://localhost:8123/req"
	req.Challenge = sso_url

	pamLog("NewRequest...")
	r, _ := http.NewRequest("POST", sso_url, nil)
	pamLog("Client...")
	client := &http.Client{}
	pamLog("Client.Do...")
	_, _ = client.Do(r)
	pamLog("PostForm...")
	_, _ = http.PostForm(sso_url, url.Values{"user": {user}})
	pamLog("Done...")
/*
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}
*/
	return nil
}

func main() {

}
