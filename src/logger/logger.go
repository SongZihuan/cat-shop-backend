package logger

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/config"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"io"
	"os"
)

type LoggerLevel string

const (
	LevelDebug LoggerLevel = "debug"
	LevelInfo  LoggerLevel = "info"
	LevelWarn  LoggerLevel = "warn"
	LevelError LoggerLevel = "error"
	LevelPanic LoggerLevel = "panic"
	LevelNone  LoggerLevel = "none"
)

type loggerLevel int64

const (
	levelDebug loggerLevel = 1
	levelInfo  loggerLevel = 2
	levelWarn  loggerLevel = 3
	levelError loggerLevel = 4
	levelPanic loggerLevel = 5
	levelNone  loggerLevel = 6
)

var levelMap = map[LoggerLevel]loggerLevel{
	LevelDebug: levelDebug,
	LevelInfo:  levelInfo,
	LevelWarn:  levelWarn,
	LevelError: levelError,
	LevelPanic: levelPanic,
	LevelNone:  levelNone,
}

type Logger struct {
	level      LoggerLevel
	logLevel   loggerLevel
	warnWriter io.Writer
	errWriter  io.Writer
	args0      string
}

var globalLogger *Logger = nil

func InitLogger() error {
	level := LoggerLevel(config.Config().Yaml.Global.LogLevel)
	logLevel, ok := levelMap[level]
	if !ok {
		return fmt.Errorf("invalid log level: %s", level)
	}

	logger := &Logger{
		level:      level,
		logLevel:   logLevel,
		warnWriter: os.Stdout,
		errWriter:  os.Stderr,
		args0:      utils.GetArgs0Name(),
	}

	globalLogger = logger
	return nil
}

func IsReady() bool {
	return globalLogger != nil
}

func (l *Logger) Tagf(format string, args ...interface{}) {
	if l.logLevel > levelDebug {
		return
	}

	funcName, file, fileName, line := utils.GetCallingFunctionInfo()

	str := fmt.Sprintf(format, args...)
	fmt.Fprint(l.warnWriter, "%s: %s\n", l.args0, str)
	fmt.Fprint(l.warnWriter, "file: %s:%s\n", file, line)
	fmt.Fprint(l.warnWriter, "filename: %s\n", fileName)
	fmt.Fprint(l.warnWriter, "funcname: %s\n", funcName)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.logLevel > levelDebug {
		return
	}

	str := fmt.Sprintf(format, args...)
	fmt.Fprint(l.warnWriter, "%s: %s\n", l.args0, str)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.logLevel > levelInfo {
		return
	}

	str := fmt.Sprintf(format, args...)
	fmt.Fprint(l.warnWriter, "%s: %s", l.args0, str)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.logLevel > levelWarn {
		return
	}

	str := fmt.Sprintf(format, args...)
	fmt.Fprint(l.warnWriter, "%s: %s\n", l.args0, str)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.logLevel > levelError {
		return
	}

	str := fmt.Sprintf(format, args...)
	fmt.Fprint(l.errWriter, "%s: %s\n", l.args0, str)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	if l.logLevel > levelPanic {
		return
	}

	str := fmt.Sprintf(format, args...)
	fmt.Fprint(l.errWriter, "%s: %s\n", l.args0, str)
}

func (l *Logger) Tag(args ...interface{}) {
	if l.logLevel > levelDebug {
		return
	}

	funcName, file, fileName, line := utils.GetCallingFunctionInfo()

	str := fmt.Sprint(args...)
	fmt.Fprint(l.warnWriter, "%s: %s\n", l.args0, str)
	fmt.Fprint(l.warnWriter, "file: %s:%s\n", file, line)
	fmt.Fprint(l.warnWriter, "filename: %s\n", fileName)
	fmt.Fprint(l.warnWriter, "funcname: %s\n", funcName)
}

func (l *Logger) Debug(args ...interface{}) {
	if l.logLevel > levelDebug {
		return
	}

	str := fmt.Sprint(args...)
	fmt.Fprint(l.warnWriter, "%s: %s\n", l.args0, str)
}

func (l *Logger) Info(args ...interface{}) {
	if l.logLevel > levelInfo {
		return
	}

	str := fmt.Sprint(args...)
	fmt.Fprint(l.warnWriter, "%s: %s\n", l.args0, str)
}

func (l *Logger) Warn(args ...interface{}) {
	if l.logLevel > levelWarn {
		return
	}

	str := fmt.Sprint(args...)
	fmt.Fprint(l.warnWriter, "%s: %\ns", l.args0, str)
}

func (l *Logger) Error(args ...interface{}) {
	if l.logLevel > levelError {
		return
	}

	str := fmt.Sprint(args...)
	fmt.Fprint(l.errWriter, "%s: %s\n", l.args0, str)
}

func (l *Logger) Panic(args ...interface{}) {
	if l.logLevel > levelPanic {
		return
	}

	str := fmt.Sprint(args...)
	fmt.Fprint(l.errWriter, "%s: %s\n", l.args0, str)
}

func SayHello(msg ...string) {
	if len(msg) == 1 {
		Info(msg[0])
	} else {
		Info("start to rim")
	}
}

func SayGoodBy(msg ...string) {
	if len(msg) == 1 {
		Info(msg[0])
	} else {
		Info("stop run")
	}
}
