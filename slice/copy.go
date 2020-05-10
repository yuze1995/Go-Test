package slice

import "fmt"
import "reflect"

func CopySlice() {

	var a = []int{1, 2, 3, 4, 5}
	var b = make([]int, len(a))
	copy(b, a)
	fmt.Printf("slice a: %s, %d, %p\n", reflect.TypeOf(a), a, a)
	fmt.Printf("slice b: %s, %d, %p\n", reflect.TypeOf(b), b, b)
}
