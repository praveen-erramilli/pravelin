package logging

import (
	"bytes"
	"log"
	"os"
)

var (
	buf      bytes.Buffer
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func init() {

	InfoLog = log.New(&buf, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLog.SetOutput(os.Stdout)

	ErrorLog = log.New(&buf, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog.SetOutput(os.Stdout)
}
