package helper

import (
	"log"
	"strings"

	"github.com/comail/colog"
)

func InitLog(level string) {
	switch strings.ToUpper(level) {
	case "TRACE":
		colog.SetMinLevel(colog.LTrace)
	case "DEBUG":
		colog.SetMinLevel(colog.LDebug)
	case "INFO":
		colog.SetMinLevel(colog.LInfo)
	case "WARNING":
		colog.SetMinLevel(colog.LWarning)
	case "ERROR":
		colog.SetMinLevel(colog.LError)
	}
	colog.SetDefaultLevel(colog.LInfo)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
	log.Printf("")
}

func TraceLog(v ...interface{}) {
	log.Printf("trace: %v", v)
}

func DebugLog(v ...interface{}) {
	log.Printf("debug: %v", v)
}

func InfoLog(v ...interface{}) {
	log.Printf("info: %+v", &v)
}

func WarnLog(v ...interface{}) {
	log.Printf("warn: %v", v)
}

func ErrorLog(v ...interface{}) {
	log.Printf("error: %v", v)
}
