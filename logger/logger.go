package logger

import (
    "log"
)

type Logger struct {
    env string
}

func New(env string) *Logger {
    return &Logger{env: env}
}

func (l *Logger) Info(v ...interface{})  { log.Println("[INFO]", v) }
func (l *Logger) Infof(format string, v ...interface{}) { log.Printf("[INFO] "+format, v...) }
func (l *Logger) Error(v ...interface{}) { log.Println("[ERROR]", v) }
func (l *Logger) Errorf(format string, v ...interface{}) { log.Printf("[ERROR] "+format, v...) }
func (l *Logger) Fatalf(format string, v ...interface{}) { log.Fatalf("[FATAL] "+format, v...) }
func (l *Logger) Fatal(v ...interface{}) { log.Fatal("[FATAL]", v) }
