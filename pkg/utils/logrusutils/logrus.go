package logrusutils

import (
	"os"

	"github.com/sirupsen/logrus"
)

func New() *logrus.Logger {
	log := logrus.New()

	log.SetOutput(os.Stdout)

	return log
}
