package utilities

import (
	// "bytes"
	"log"
	"os"
)

var (
	InfoLog    *log.Logger = log.New(os.Stdout, "[LOG] ", log.LstdFlags)
	ErrorLog   *log.Logger = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
	WarningLog *log.Logger = log.New(os.Stdout, "[WARNING] ", log.LstdFlags)
)

func PrintInfo(args ...any) {
	InfoLog.Print(":")
	InfoLog.Println(args...)
}

func PrintError(args ...any) {
	InfoLog.Print(":")
	ErrorLog.Println(args...)
}

func PrintWarning(args ...any) {
	InfoLog.Print(":")
	WarningLog.Println(args...)
}
