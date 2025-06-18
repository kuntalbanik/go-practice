package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 22.5
	num2 := int(num)
	fmt.Println(num2)

	fmt.Printf("Num type is %T\n", num2)
	newstr1 := strconv.Itoa(num2)
	fmt.Println(newstr1)
	fmt.Printf("Data type is %T\n", newstr1)

	//
	//

	newstr := "hello"
	num3, _ := strconv.Atoi(newstr)
	fmt.Println(num3)
	fmt.Printf("Data type is %T\n", num3)

	//
	//

	numFloat := "3.142"
	num4, _ := strconv.ParseFloat(numFloat, 64)
	fmt.Println(num4)
	fmt.Printf("Data type is %T\n", num4)

}
