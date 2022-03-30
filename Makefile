MODULE := pam_websso

ALL: module

deps:
	GOPATH=${PWD}/.go go get -d

module: deps
	GOPATH=${PWD}/.go go build -buildmode=c-shared -o ${MODULE}.so

install: module
	sudo cp ${MODULE}.so /lib/x86_64-linux-gnu/security/

clean:
	go clean
	-rm -f ${MODULE}.so ${MODULE}.h
	-rm -rf .go/

test: install
	pamtester sshd worker authenticate