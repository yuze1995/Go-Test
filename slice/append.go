package slice

import "fmt"
import "reflect"

func AppendSlice1() {
	var slice0 = make([]int, 0)
	fmt.Println("start slice", slice0)
	for i := 0; i < 10; i++ {
		slice0 = append(slice0, i+1)
		fmt.Printf("index: %d, len: %d, cap: %d, pointer: %p, content: %v \n", i, len(slice0), cap(slice0), slice0, slice0)
	}
}

func AppendSlice2() {
	var slice0 = make([]int, 0)
	fmt.Printf("%s, %d, %p\n", reflect.TypeOf(slice0), slice0, slice0)
	slice1 := append(slice0, 1)
	fmt.Printf("%s, %d, %p\n", reflect.TypeOf(slice1), slice1, slice1)
	slice2 := append(slice1, 1)
	fmt.Printf("%s, %d, %p\n", reflect.TypeOf(slice2), slice2, slice2)
	slice3 := append(slice2, 1)
	fmt.Printf("%s, %d, %p\n", reflect.TypeOf(slice3), slice3, slice3)
	slice4 := append(slice3, 1)
	fmt.Printf("%s, %d, %p\n", reflect.TypeOf(slice4), slice4, slice4)
}

func CreateSliceFromSlice() {
	a := make([]int, 10, 20)
	b := a[5:]
	fmt.Printf("slice a: %s, %d, %p\n", reflect.TypeOf(a), a, a)
	fmt.Printf("slice b: %s, %d, %p\n", reflect.TypeOf(b), b, b)
	b = append(b, 0)
	fmt.Println("add one element to slice b")
	fmt.Printf("slice a: %s, %d, %p\n", reflect.TypeOf(a), a, a)
	fmt.Printf("slice b: %s, %d, %p\n", reflect.TypeOf(b), b, b)
}

func PassSlice() {
	var a = []int{1, 2, 3, 4, 5}
	var b = a
	doAppend(a)
	fmt.Printf("slice a: %s, %d, %p\n", reflect.TypeOf(a), a, a)
	fmt.Printf("slice b: %s, %d, %p\n", reflect.TypeOf(b), b, b)
	doAppend(a[0:2])
	fmt.Printf("slice a: %s, %d, %p\n", reflect.TypeOf(a), a, a)
	fmt.Printf("slice b: %s, %d, %p\n", reflect.TypeOf(b), b, b)
}

func doAppend(a1 []int) {
	fmt.Printf("slice a1 before append: %s, %d, %p\n", reflect.TypeOf(a1), a1, a1)
	b1 := append(a1, 0)
	fmt.Printf("slice a1 after append: %s, %d, %p\n", reflect.TypeOf(a1), a1, a1)
	fmt.Printf("slice b1 after append: %s, %d, %p\n", reflect.TypeOf(b1), b1, b1)
}
