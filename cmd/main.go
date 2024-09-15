package main

import (
	"io"
	"log"
	"os"
)

func main() {
	loggingSettiings("./log/output.log")
	s := server.NewServer()
	s.Run()
}

func loggingSettings(filename stirng) {
	logfile, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llogfile)
	log.SetOutput(multiLogFile)
}
