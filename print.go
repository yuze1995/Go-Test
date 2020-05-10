package main

import "fmt"
import "Go-Test/print"
import log "github.com/sirupsen/logrus"

func Print(){
	fmt.Println("Hi")
	PrintLine()

	log.Info("HiHi")
	PrintLine()

	print.Hello()
	PrintLine()
}