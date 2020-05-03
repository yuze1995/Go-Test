package array

import "fmt"

var array2 = [2]int{1, 2}

func PassValueType() {
	fmt.Println("Pass ValueType:")
	fmt.Printf("address: %p, value: %v \n", &array2, array2)
	valueType(array2)
}

func PassReferenceType() {
	fmt.Println("Pass ReferenceType:")
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
	for i := range array {
		array[i] += 1
	}
	fmt.Println("old array add 1")
	fmt.Printf("old array: %v, ", array2)
	fmt.Printf("new array: %v \n", array)
}
