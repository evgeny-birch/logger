package logger

import (
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	log.SetOutput(os.Stdout)
}

func Info(format string, v ...interface{}) {
	log.Printf("[INFO] "+format, v...)
}

func Warning(format string, v ...interface{}) {
	log.Printf("[WARNING] "+format, v...)
}

func Error(format string, v ...interface{}) {
	log.Printf("[ERROR] "+format, v...)
}
