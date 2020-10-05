package main

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"github.com/newrelic/go-agent/v3/integrations/logcontext/nrlogrusplugin"
)

type Writer struct{}

func (w Writer) Write(p []byte) (n int, err error) {
	return 0, nil
}

var (
	logger *logrus.Logger
)

func main() {
	pathMap := lfshook.PathMap{
		logrus.PanicLevel: "error.log",
		logrus.ErrorLevel: "error.log",
		logrus.FatalLevel: "error.log",
		logrus.WarnLevel:  "error.log",
		logrus.InfoLevel:  "info.log",
		logrus.DebugLevel: "info.log",
	}

	logger = logrus.New()
	logger.Hooks.Add(lfshook.NewHook(
		pathMap,
		nrlogrusplugin.ContextFormatter{},
	))

	logger.SetOutput(Writer{})

	logger.Infoln("test info baru lagi")
	logger.Errorln("test error baru lagi")
}
