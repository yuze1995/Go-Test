package slice

import "fmt"
import "reflect"

func Element() {
	var array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := array[1:3]
	fmt.Println(array)
	fmt.Println(reflect.TypeOf(array))
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println()
}

func MakeSlice() {
	var slice3 = make([]int, 3)
	var slice10 = make([]int, 3, 10)
	fmt.Println(slice3, slice10)
	fmt.Println(reflect.TypeOf(slice3), reflect.TypeOf(slice10))
	fmt.Println("slice3 now length:", len(slice3), ", max length:", cap(slice3))
	fmt.Println("slice10 now length:", len(slice10), ", max length:", cap(slice10))
}

func AppendSlice() {
	var slice0 = make([]int, 0)
	fmt.Println("start slice", slice0)
	for i := 0; i < 10; i++ {
		slice0 = append(slice0, i+1)
		fmt.Printf("index: %d, len: %d, cap: %d, pointer: %p, content: %v \n", i, len(slice0), cap(slice0), slice0, slice0)
	}
}
