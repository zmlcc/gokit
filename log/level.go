package log

import (
	"errors"
)

type Level uint8

const (
	ErrorLevel Level = iota
	WarnLevel
	InfoLevel
	DebugLevel
)

var ErrLevelMatch = errors.New("Log Level Not Match")

// Convert the Level to a string.
func (level Level) String() string {
	switch level {
	case DebugLevel:
		return "DEBU"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERRO"
	default:
		return "NONE"
	}

}

type LevelLogger struct {
	ctxLog *Context
	level  Level
}

//NewLevel creates a new leveled logger,
func NewLevel(log Logger, lvl Level) *LevelLogger {
	return &LevelLogger{
		ctxLog: NewContext(log),
		level:  lvl,
	}
}

func (lg *LevelLogger) Log(keyvals ...interface{}) error {
	return lg.ctxLog.Log(keyvals...)
}

func (lg *LevelLogger) With(keyvals ...interface{}) *LevelLogger {
	return &LevelLogger{
		ctxLog: lg.ctxLog.With(keyvals...),
		level:  lg.level,
	}
	// lg.ctxLog = lg.ctxLog.With(keyvals...)
	// return lg
}

func (lg *LevelLogger) WithLevel(lvl Level) *LevelLogger {
	return &LevelLogger{
		ctxLog: lg.ctxLog,
		level:  lvl,
	}
}

func (lg *LevelLogger) levelLog(lvl Level, keyvals ...interface{}) error {
	if lg.level < lvl {
		return ErrLevelMatch
	}
	return lg.ctxLog.WithPrefix("lvl", lvl.String()).Log(keyvals...)
}

func (lg *LevelLogger) Debug(kvs ...interface{}) error {
	return lg.levelLog(DebugLevel, kvs...)
}

func (lg *LevelLogger) Info(kvs ...interface{}) error {
	return lg.levelLog(InfoLevel, kvs...)
}

func (lg *LevelLogger) Warn(kvs ...interface{}) error {
	return lg.levelLog(WarnLevel, kvs...)
}

func (lg *LevelLogger) Error(kvs ...interface{}) error {
	return lg.levelLog(ErrorLevel, kvs...)
}

