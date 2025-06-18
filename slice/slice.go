package slice

import "fmt"

func SliceCustom() {
	numbers := make([]int, 20)

	for i := range numbers {
		numbers[i] = i
	}
	//
	new1 := numbers[0:5]  // slicing from slice
	new2 := numbers[6:15] // slicing from slice
	//
	//
	fmt.Println(numbers)
	fmt.Println(new1)
	fmt.Println(new2)
}

func SliceModifier(n []int) []int {
	for i := range n {
		n[i] = i
	}
	n = append(n, 50)

	// copy slice
	nn := make([]int, len(n))
	copy(nn, n)

	nn = append(nn, 100)

	return nn

}
