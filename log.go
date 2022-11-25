package nutil

import (
	"fmt"
	"io"
	"log"
	"os"
)

type BasicLogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
	printLevel  int
	flag        int
}

const (
	//1, 2, 4, 8, 16
	logLevelError = 1 << iota
	logLevelWarn
	logLevelInfo
	logLevelDebug
)
const (
	LogPrintError = logLevelError
	LogPrintWarn  = LogPrintError | logLevelWarn
	LogPrintInfo  = LogPrintWarn | logLevelInfo
	LogPrintDebug = LogPrintInfo | logLevelDebug
)

//like log package flags
const (
	Ldate         = log.Ldate
	Ltime         = log.Ltime
	Lmicroseconds = log.Lmicroseconds
	Llongfile     = log.Llongfile
	Lshortfile    = log.Lshortfile
	LUTC          = log.LUTC
	Lmsgprefix    = log.Lmsgprefix
	LstdFlags     = log.LstdFlags
)

func createDefaultLogger(prefix string) *log.Logger {

	prefixStr := fmt.Sprintf("[%-5s] ", prefix)
	newLogger := log.New(os.Stdout, prefixStr, LstdFlags|Lmsgprefix)
	return newLogger
}

func CreateBasicLogger() *BasicLogger {

	ret := &BasicLogger{
		infoLogger:  createDefaultLogger("INFO"),
		debugLogger: createDefaultLogger("DEBUG"),
		errorLogger: createDefaultLogger("ERROR"),
		warnLogger:  createDefaultLogger("WARN"),
		printLevel:  LogPrintInfo,
		flag:        LstdFlags | Lmsgprefix,
	}

	return ret
}

func (r *BasicLogger) isPrintAble(logLevel int) bool {

	if r.printLevel&logLevel != 0 {
		return true
	}
	return false
}

func (r *BasicLogger) Info(str string, args ...interface{}) {
	if r.isPrintAble(logLevelInfo) {
		r.infoLogger.Printf(str, args...)
	}
}

func (r *BasicLogger) Warn(str string, args ...interface{}) {
	if r.isPrintAble(logLevelWarn) {
		r.warnLogger.Printf(str, args...)
	}
}

func (r *BasicLogger) Debug(str string, args ...interface{}) {
	if r.isPrintAble(logLevelDebug) {
		r.debugLogger.Printf(str, args...)
	}
}

func (r *BasicLogger) Error(str string, args ...interface{}) {
	if r.isPrintAble(logLevelError) {
		r.errorLogger.Printf(str, args...)
	}
}

func (r *BasicLogger) SetLogLevel(logLevel int) {
	if LogPrintDebug&logLevel != 0 {
		r.printLevel = logLevel
	}
}

func (r *BasicLogger) SetWriter(writer io.Writer) {
	r.debugLogger.SetOutput(writer)
	r.infoLogger.SetOutput(writer)
	r.warnLogger.SetOutput(writer)
	r.errorLogger.SetOutput(writer)
}

func (r *BasicLogger) SetFlags(flag int) {
	r.debugLogger.SetFlags(flag)
	r.infoLogger.SetFlags(flag)
	r.warnLogger.SetFlags(flag)
	r.errorLogger.SetFlags(flag)
	r.flag = flag
}

func (r *BasicLogger) Flags() int {
	return r.flag
}
