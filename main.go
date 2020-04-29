package main

import "fmt"

import "Go-Test/print"
import "Go-Test/array"

import log "github.com/sirupsen/logrus"

func main() {

	fmt.Println("Hi")
	printLine()
	log.Info("HiHi")
	printLine()
	print.Hello()
	printLine()
	array.Element3("a")
	printLine()
	array.Element4("b")
	printLine()
	array.PassValueType()
	printLine()
	array.PassReferenceType()
	printLine()
}

func printLine() {
	fmt.Println("------------------------------------")
}
