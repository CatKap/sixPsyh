package loger

import (
    "log"
)

type Loger struct {
    env string
}

func New(env string) *Loger {
    return &Loger{env: env}
}

func (l *Loger) Info(v ...interface{})  { log.Println("[INFO]", v) }
func (l *Loger) Infof(format string, v ...interface{}) { log.Printf("[INFO] "+format, v...) }
func (l *Loger) Error(v ...interface{}) { log.Println("[ERROR]", v) }
func (l *Loger) Errorf(format string, v ...interface{}) { log.Printf("[ERROR] "+format, v...) }
func (l *Loger) Fatalf(format string, v ...interface{}) { log.Fatalf("[FATAL] "+format, v...) }
func (l *Loger) Fatal(v ...interface{}) { log.Fatal("[FATAL]", v) }
