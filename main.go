package main

import "fmt"

import "Go-Test/print"
import "Go-Test/array"
import "Go-Test/slice"

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

	slice.Element()
	printLine()

	slice.MakeSlice()
	printLine()

	slice.AppendSlice()
	printLine()
}

func printLine() {
	fmt.Println("------------------------------------")
}
