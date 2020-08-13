/*
@Time    : 2020/8/1
@Author  : Wangcq
@File    : log.go
@Software: GoLand
*/

package logs

import (
	"github.com/keepeye/logrus-filename"
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func init() {
	filenameHook := filename.NewHook()
	filenameHook.Field = "line"
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetOutput(os.Stdout)
	Log.AddHook(filenameHook)
	if env := os.Getenv("env"); env == "dev" {
		Log.SetLevel(logrus.DebugLevel)
	} else {
		Log.SetLevel(logrus.InfoLevel)
	}

}
