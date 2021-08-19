package testable

import (
	"io/ioutil"
	"time"
)

func Write(name string) error {

	ft := time.Now().Format(time.RFC3339)

	return ioutil.WriteFile(name, []byte(ft), 0644)
}
