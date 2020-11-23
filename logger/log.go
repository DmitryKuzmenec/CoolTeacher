package logger

import (
	"github.com/DmitryKuzmenec/CoolTeacher/config"
	"github.com/sirupsen/logrus"
)

// Logger - wrapper of logrus
type Logger struct {
	*logrus.Logger
}

// Init - constructor of Logger
func Init(conf *config.Config) (*Logger, error) {
	log := logrus.New()

	level, err := logrus.ParseLevel(conf.Log.Level)
	if err != nil {
		return nil, err
	}
	log.SetLevel(level)

	// set formatting
	formatter := &logrus.TextFormatter{
		DisableColors: false,
		ForceQuote:    true,
		FullTimestamp: true,
	}

	log.SetFormatter(formatter)
	return &Logger{log}, nil
}
