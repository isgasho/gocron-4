package gocron

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

var (
	logger Logger
)

func initLogger(l Logger) {
	logger = l
}

func debugf(format string, args ...interface{}) {
	if logger != nil {
		logger.Debugf(format, args)
	}
}

func infof(format string, args ...interface{}) {
	if logger != nil {
		logger.Infof(format, args)
	}
}

func errorf(format string, args ...interface{}) {
	if logger != nil {
		logger.Errorf(format, args)
	}
}
