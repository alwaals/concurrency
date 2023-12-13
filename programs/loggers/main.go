package main

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	f, err := os.OpenFile("loggers.text", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("Error", err)
	}
	log.SetOutput(f)
	log.Printf("Error received from handler\n")

	WarningLogger = log.New(f, "WARNING:", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(f, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(f, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	WarningLogger.Println("This is just a warning")
	InfoLogger.Println("This is just a testing information")
	ErrorLogger.Println("Got a serious error")
}
