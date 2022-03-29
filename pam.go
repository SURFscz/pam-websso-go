// +build darwin linux

package main

import (
	"unsafe"
)

/*
#cgo LDFLAGS: -lpam -fPIC
#include <security/pam_appl.h>
#include <stdlib.h>

int pam_prompt(pam_handle_t *pamh, int style, char **response, const char *fmt);
*/
import "C"

func init() {
	if !disablePtrace() {
		pamLog("unable to disable ptrace")
	}
}

//export pam_sm_authenticate
func pam_sm_authenticate(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	var tempBuf *C.char
	defer C.free(unsafe.Pointer(tempBuf))

	// Load config from argv
        config := Config{}
        config.LoadFromFile(argv)

	user := "martin"
	req := Req{}
	req.startReq(user)

	C.pam_prompt(pamh, C.PAM_PROMPT_ECHO_ON, nil, C.CString(req.Challenge))
	//C.pam_prompt(pamh, C.PAM_PROMPT_ECHO_ON, nil, C.CString(config.Sso_server))
	//C.pam_prompt(pamh, C.PAM_PROMPT_ECHO_ON, &tempBuf, C.CString("Press enter (Go pam_websso)"))

	return C.PAM_SUCCESS
}

//export pam_sm_setcred
func pam_sm_setcred(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	return C.PAM_SUCCESS
}
