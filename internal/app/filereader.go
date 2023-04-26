package app

import (
	"io/ioutil"
)

func ReadStartFile() (error, string) {
	FileString, err := ioutil.ReadFile("configs/texts/StartMessage.txt")
	if err != nil {
		return err, ""
	}
	return nil, string(FileString)
}
