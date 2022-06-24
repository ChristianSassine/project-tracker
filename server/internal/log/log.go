package log

import (
	"log"
	"os"
)

var (
	InfoLog    *log.Logger = log.New(os.Stdout, "[LOG] ", log.LstdFlags)
	ErrorLog   *log.Logger = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
	WarningLog *log.Logger = log.New(os.Stdout, "[WARNING] ", log.LstdFlags)
)

const suffix = ":"

func PrintInfo(args ...any) {
	args = append([]any{suffix}, args...)
	InfoLog.Println(args...)
}

func PrintError(args ...any) {
	args = append([]any{suffix}, args...)
	ErrorLog.Println(args...)
}

func PrintWarning(args ...any) {
	args = append([]any{suffix}, args...)
	WarningLog.Println(args...)
}
