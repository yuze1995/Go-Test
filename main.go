package main

import "fmt"

import "go_test/print"
import "go_test/array"

import log "github.com/sirupsen/logrus"

func main() {

	fmt.Println("Hi")
	log.Info("HiHi")
	print.Hello()
	array.Element3("a")
	array.Element4("b")
}
