package anothertestable

import (
	"io/ioutil"
	"time"
)

// "strings"
type tester interface {
	//dataGenerate() string
	//timeGenerate() time.Time
	dataGenerate()
	timeGenerate()
}
type writeData struct {
	//data dataGenerate()
	//time timeGenerate()
	data string
	time time.Time
}

func (w *writeData) dataGenerate(information string) (string, error) {

	return w.data, nil
}
func (w *writeData) timeGenerate() (time.Time, error) {
	return w.time, nil

}
func write(name string) error {
	var dg writeData
	dg.dataGenerate("data")
	ft := time.Now().Format(time.RFC3339)

	//ft := time.Now().Format(time.RFC3339)

	return ioutil.WriteFile(name, []byte(ft), 0644)
}
