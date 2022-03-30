MODULE := pam_websso

ALL: module

deps:
	GOPATH=${HOME}/.go go get -d

module: deps
	GOPATH=${HOME}/.go go build -buildmode=c-shared -o ${MODULE}.so

install: module
	sudo sed -i '/'${MODULE}'/d' /etc/pam.d/sshd
	echo "auth required ${MODULE}.so ${PWD}/config.yaml" | sudo tee -a /etc/pam.d/sshd
	sudo cp ${MODULE}.so /lib/x86_64-linux-gnu/security/

clean:
	go clean
	-rm -f ${MODULE}.so ${MODULE}.h
	-rm -rf .go/

test: install
	pamtester sshd worker authenticate