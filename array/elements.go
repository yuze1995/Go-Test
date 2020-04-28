package array

import "fmt"
import "strconv"

func Element3(prev string) {

	seq := [3]int{1, 2, 3}

	for i := range seq {
		print(prev, seq[i])
	}
	fmt.Println("")
}

func Element4(prev string) {
	var seq = [...]int{1, 2, 3, 4}

	for i := 0; i < len(seq); i++ {
		print(prev, seq[i])
	}
	fmt.Println("")
}

func print(prev string, count int) {
	fmt.Print(prev + strconv.Itoa(count) + " ")
}
