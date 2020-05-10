package slice

import "fmt"
import "reflect"

func DeleteSlice() {

	var a = []int{1, 2, 3, 4, 5}
	fmt.Printf("slice a: %s, %d, %p\n", reflect.TypeOf(a), a, a)
	var b = append(a[:1], a[1+1:]...)
	fmt.Printf("slice b: %s, %d, %p\n", reflect.TypeOf(b), b, b)
}
