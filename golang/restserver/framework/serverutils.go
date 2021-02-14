package framework

import (
	"github.com/sirupsen/logrus"
	"net/url"
	"regexp"
	"runtime"
	"time"
)

func CheckUrl(str string) bool {
	_, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}

	u, err := url.Parse(str)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)

	pc, _, _, _ := runtime.Caller(1)

	funcObj := runtime.FuncForPC(pc)

	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

	logrus.Infof("%s took %s", name, elapsed)
}