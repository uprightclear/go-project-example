package log

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
)

var l map[Log]Logger

type Log string

const (
	AccessLog Log = "accessLog"
	AppLog    Log = "appLog"
)

type Logger struct {
	logger *zap.SugaredLogger
}

type Level string

const (
	Debug  Level = "debug"
	Info   Level = "info"
	Warn   Level = "warn"
	Error  Level = "error"
	DPanic Level = "dpanic"
	Panic  Level = "panic"
	Fatal  Level = "fatal"
)

var levelMap = map[Level]zapcore.Level{
	Debug:  zapcore.DebugLevel,
	Info:   zapcore.InfoLevel,
	Warn:   zapcore.WarnLevel,
	Error:  zapcore.ErrorLevel,
	DPanic: zapcore.DPanicLevel,
	Panic:  zapcore.PanicLevel,
	Fatal:  zapcore.FatalLevel,
}

type Config struct {
	Path string
}

// Initialize init log
func Initialize(configs map[Log]Config, level Level) (err error) {
	l = make(map[Log]Logger)

	for log, config := range configs {
		if l[log], err = create(config.Path, level); err != nil {
			return
		}
	}

	return
}

// create returns logger
func create(path string, level Level) (logger Logger, err error) {
	if !strings.HasSuffix(path, ".log") {
		err = errors.New("log path must be .log file")
		return
	}

	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    500,
		MaxBackups: 200,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   true,
	})

	var lev zapcore.Level
	var ok bool
	if lev, ok = levelMap[level]; !ok {
		lev = levelMap[Info] // default level is info
	}
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(lev))

	logger = Logger{
		logger: zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar(),
	}

	return
}

func Get(log Log) Logger {
	return l[log]
}

func (l Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l Logger) DPanic(args ...interface{}) {
	l.logger.DPanic(args...)
}

func (l Logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l Logger) DPanicf(format string, args ...interface{}) {
	l.logger.DPanicf(format, args...)
}

func (l Logger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}