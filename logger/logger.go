package logger

import (
	"time"

	"github.com/sirupsen/logrus"
	formatter "github.com/tim-ywliu/nested-logrus-formatter"
)

var (
	log     *logrus.Logger
	NgapLog *logrus.Entry
)

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	}

	NgapLog = log.WithFields(logrus.Fields{"component": "LIB", "category": "NGAP"})
}

func GetLogger() *logrus.Logger {
	return log
}

func SetLogLevel(level logrus.Level) {
	NgapLog.Infoln("set log level :", level)
	log.SetLevel(level)
}

func SetReportCaller(enable bool) {
	NgapLog.Infoln("set report call :", enable)
	log.SetReportCaller(enable)
}
