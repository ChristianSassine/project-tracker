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
