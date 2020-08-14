package utils

import (
	"github.com/sirupsen/logrus"
	"runtime/debug"
)

func Recovery() {
	if r := recover(); r != nil {
		logrus.Errorf("catch panic! trace:\n %s", debug.Stack())
	}
}
