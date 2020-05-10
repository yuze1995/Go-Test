package main

import "Go-Test/array"

func Array() {
	array.Element3("a")
	PrintLine()

	array.Element4("b")
	PrintLine()

	array.PassValueType()
	PrintLine()

	array.PassReferenceType()
	PrintLine()
}
