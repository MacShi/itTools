package utils
import (
"fmt"
"log"
"runtime"
)

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var logLevel = INFO

func SetLogLevel(level int) {
	logLevel = level
}

func logWithLevel(level int, format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(2) // 获取调用者的信息
	logHeader := fmt.Sprintf("[%s:%d] ", file, line)

	if level >= logLevel {
		switch level {
		case DEBUG:
			log.Printf(logHeader+"[DEBUG] "+format, args...)
		case INFO:
			log.Printf(logHeader+"[INFO] "+format, args...)
		case WARNING:
			log.Printf(logHeader+"[WARNING] "+format, args...)
		case ERROR:
			log.Printf(logHeader+"[ERROR] "+format, args...)
		case FATAL:
			log.Fatalf(logHeader+"[FATAL] "+format, args...)
		}
	}
}

func Debug(format string, args ...interface{}) {
	logWithLevel(DEBUG, format, args...)
}

func Info(format string, args ...interface{}) {
	logWithLevel(INFO, format, args...)
}

func Warning(format string, args ...interface{}) {
	logWithLevel(WARNING, format, args...)
}

func Error(format string, args ...interface{}) {
	logWithLevel(ERROR, format, args...)
}

func Fatal(format string, args ...interface{}) {
	logWithLevel(FATAL, format, args...)
}
