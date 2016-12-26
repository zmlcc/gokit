package log

import (
    "os"
)

var (
    // std = NewLevel(NewLogfmtLogger(NewSyncWriter(os.Stdout)), DebugLevel)
    std = NewLevel(NewLogfmtLogger(os.Stdout), DebugLevel)
)


func With(kvs ...interface{}) *LevelLogger {
    return std.With(kvs...)
}

func WithLevel(lvl Level) *LevelLogger{
    return std.WithLevel(lvl)
}

func Debug(kvs ...interface{}) error {
    return std.Debug(kvs...)
}


func Info(kvs ...interface{}) error {
    return std.Info(kvs...)
}

func Warn(kvs ...interface{}) error {
    return std.Warn(kvs...)
}

func Error(kvs ...interface{}) error {
    return std.Error(kvs...)
}