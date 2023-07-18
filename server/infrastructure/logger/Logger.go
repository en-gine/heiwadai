package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func (l LogLevel) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "Unknown"
	}
}

var LogConfig = &lumberjack.Logger{
	Filename:   "/server/log/app.log", // ファイル名
	MaxSize:    500,                   // ローテーションするファイルサイズ(megabytes)
	MaxBackups: 3,                     // 保持する古いログの最大ファイル数
	MaxAge:     365,                   // 古いログを保持する日数
	LocalTime:  true,                  // バックアップファイルの時刻フォーマットをサーバローカル時間指定
	Compress:   true,                  // ローテーションされたファイルのgzip圧縮
}

func Log(level LogLevel, message string) {
	mode := os.Getenv("ENV_MODE")
	if mode == "production" && level == LevelDebug {
		return
	}

	_, file, line, _ := runtime.Caller(2)
	file = strings.Replace(file, os.Getenv("GOPATH")+"/src/", "", 1)
	now := time.Now()
	log.SetOutput(LogConfig)

	log.Printf("%s [%s] %s:%d %s\n", now.Format("2006-01-02 15:04:05"), level.String(), file, line, message)
}

func Error(message string) {
	Log(LevelWarn, message)
}

func Errorf(format string, a ...interface{}) {
	Log(LevelError, fmt.Sprintf(format, a...))
}

func Info(message string) {
	Log(LevelInfo, message)
}

func Infof(format string, a ...interface{}) {
	Log(LevelInfo, fmt.Sprintf(format, a...))
}

func Debugf(format string, a ...interface{}) {
	Log(LevelDebug, fmt.Sprintf(format, a...))
}

func Debug(message string) {
	Log(LevelDebug, message)
}

func Fatal(message string) {
	Log(LevelFatal, message)
}

func Fatalf(format string, a ...interface{}) {
	Log(LevelFatal, fmt.Sprintf(format, a...))
}

func Warn(message string) {
	Log(LevelWarn, message)
}

func Warnf(format string, a ...interface{}) {
	Log(LevelWarn, fmt.Sprintf(format, a...))
}
