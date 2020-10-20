package log

import (
	"io/ioutil"
	"log"
	"os"
)

func Discard() *log.Logger {
	return log.New(ioutil.Discard, "", 0)
}

func StdErr() *log.Logger {
	return log.New(os.Stderr, "", log.LstdFlags)
}
