package array

import "fmt"

var array2 = [2]int{1, 2}

func PassValueType() {
	fmt.Println("PassValueType:")
	fmt.Printf("address: %p, value: %v \n", &array2, array2)
	valueType(array2)
}

func PassReferenceType() {
	fmt.Println("PassReferenceType:")
	fmt.Printf("address: %p, value: %v \n", &array2, array2)
	referenceType(&array2)
}

func valueType(array [2]int) {
	fmt.Printf("new address: %p, new content: ", &array)
	fmt.Println(array)
}

func referenceType(array *[2]int) {
	fmt.Printf("new address: %p, new content: ", &array)
	fmt.Print(array)
	fmt.Printf(", new content format to pointer: %p \n", array)
}
