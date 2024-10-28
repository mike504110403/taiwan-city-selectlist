package main

import (
	"fmt"

	mlog "github.com/mike504110403/goutils/log"

	"testing"
)

func TestSendGet(t *testing.T) {
	mlog.Init(mlog.Config{
		EnvMode: "debug",
		LogType: "console",
	})
	res, err := getCountry()
	if err != nil {
		mlog.Error(err.Error())
	}
	mlog.Info(fmt.Sprintf("%+v", res))
}

func TestMain(t *testing.T) {
	mlog.Init(mlog.Config{
		EnvMode: "debug",
		LogType: "console",
	})
	main()
}
