package utilities

import (
	// "bytes"
	"log"
	"os"
)

var (
	InfoLog    *log.Logger = log.New(os.Stdout, "[LOG] ", log.LstdFlags|log.Lshortfile)
	ErrorLog   *log.Logger = log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lshortfile)
	WarningLog *log.Logger = log.New(os.Stdout, "[WARNING] ", log.LstdFlags|log.Lshortfile)
)

func PrintInfo(args ...any) {
	InfoLog.Println(args...)
}

func PrintError(args ...any) {
	ErrorLog.Println(args...)
}

func PrintWarning(args ...any) {
	InfoLog.Println(args...)
}
