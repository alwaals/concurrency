package main

import (
	"log"
	"os"
	"go.uber.org/zap"
)

func main(){
	f,_:=os.OpenFile("zapperLogger.text",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0666)
	log.SetOutput(f)
	z:=zap.L()
	z.Info("Starting the service")
	z.Debug("Debugging the message")
}