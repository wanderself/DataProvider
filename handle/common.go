package handle

import (
		"os"
	"github.com/golang/glog"
)

func Log(v ...interface{}) {
	glog.Info(v...)
}

func Err(err error) {
	if err != nil {
		Log(err)

		os.Exit(0)
	}
}
