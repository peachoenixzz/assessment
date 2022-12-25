package log

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

const InfoFile = "tmp/info.log"
const ErrorFile = "tmp/error.log"

func ErrorLog(errDetail interface{}, serv string) {
	f, err := os.OpenFile(ErrorFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Failed to create logfile" + ErrorFile)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("bye")
		}
	}(f)
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	// Output to stdout instead of the default stderr
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     false,
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// Only log the debug severity or above
	log.SetLevel(log.TraceLevel)
	//fmt.Println(err)
	log.WithFields(log.Fields{
		"service": serv,
	}).Error(errDetail)
}

func InfoLog(infoDetail interface{}, serv string) {
	f, err := os.OpenFile(InfoFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Failed to create logfile" + InfoFile)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("bye")
		}
	}(f)
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	// Output to stdout instead of the default stderr
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetFormatter(&log.JSONFormatter{})
	// Only log the debug severity or above
	log.SetLevel(log.TraceLevel)
	//fmt.Println(err)
	log.WithFields(log.Fields{
		"service": serv,
	}).Info(infoDetail)
}
