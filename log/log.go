package log

import (
	"fmt"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func Info(msg string) {
	fmt.Printf("%s[INFO] %s %s%s\n", Green, time.Now().Format("2006-01-02 15:04:05"), msg, Reset)
}

func Warn(msg string) {
	fmt.Printf("%s[WARN] %s %s%s\n", Yellow, time.Now().Format("2006-01-02 15:04:05"), msg, Reset)
}

func Error(msg string) {
	fmt.Printf("%s[ERROR] %s %s%s\n", Red, time.Now().Format("2006-01-02 15:04:05"), msg, Reset)
}
