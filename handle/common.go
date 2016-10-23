package handle

import (
	"log"
	"os"
)

func Log(v ...interface{}) {
	log.Println(v...)
}

func Err(err error) {
	if err != nil {
		Log(err)
		os.Exit(0)
	}
}
