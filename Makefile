MODULE := pam_websso

module:
	GOPATH=${PWD}/.go go build -buildmode=c-shared -o ${MODULE}.so

install:
	cp ${MODULE}.so /lib/x86_64-linux-gnu/security/

clean:
	go clean
	-rm -f ${MODULE}.so ${MODULE}.h
	-rm -rf .go/

.PHONY: module
