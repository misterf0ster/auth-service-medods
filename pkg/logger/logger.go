package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func LoggerInit() {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.DebugLevel)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

// Логирование ошибки
func LogError(context string, err error) error {
	if err != nil {
		Log.Errorf("%s: %v", context, err)
		return fmt.Errorf("%s: %w", context, err)
	}
	return nil
}

// инфо-сообщения
func LogInfo(context string) {
	Log.Infof("%s", context)
}

// debug-сообщения
func LogDebug(context string) {
	Log.Debugf("%s", context)
}

func LogFatal(context string, err error) {
	if err != nil {
		Log.Fatalf("%s: %v", context, err)
	}
}
