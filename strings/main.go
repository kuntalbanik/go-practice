package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,grape,orange,banana"
	output := strings.Split(data, ",")
	fmt.Println(output)
	fmt.Printf("%T\n", output)

	num := "one two three one two four 5 6 5 6 6 6 8 9 9 10"
	output_count := strings.Count(num, "two")
	fmt.Println(output_count)
}
