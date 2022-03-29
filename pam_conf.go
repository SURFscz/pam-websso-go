package main

import (
	"unsafe"
	"io/ioutil"
	"gopkg.in/yaml.v3"
)
/*
#cgo LDFLAGS: -lpam -fPIC
#include <security/pam_appl.h>
#include <stdlib.h>

char *string_from_argv(int, char**);
*/
import "C"


type Config struct {
	Sso_server	string	`yaml:"sso_server"`
	Token		string	`yaml:"token"`
}

// Load config from argv[0
func (c *Config) LoadFromFile(argv **C.char) error {
	path := C.string_from_argv(0, argv)
	defer C.free(unsafe.Pointer(path))

	f, err := ioutil.ReadFile(C.GoString(path))
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(f, c)
	if err != nil {
		return err
	}

	if c.Sso_server == "" {
		c.Sso_server = "http://localhost:8123"
	}
	if c.Token == "" {
		c.Sso_server = "no_token"
	}

	return nil
}

