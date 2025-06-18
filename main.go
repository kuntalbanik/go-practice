package main

import (
	"fmt"
	"go-practice/arraydata"
	"go-practice/slice"
	"strconv"
)

func main() {
	num := 22.5
	num2 := int(num)
	fmt.Println(num2)

	//
	// Number to string conversion
	//
	fmt.Printf("Num type is %T\n", num2)
	newstr1 := strconv.Itoa(num2)
	fmt.Println(newstr1)
	fmt.Printf("Data type is %T\n", newstr1)

	//
	// string to integer number - If possible to change string to number
	//
	newstr := "hello"
	num3, _ := strconv.Atoi(newstr)
	fmt.Println(num3)
	fmt.Printf("Data type is %T\n", num3)

	//
	// string to float number - If possible to change string to number
	//
	numFloat := "3.142"
	num4, _ := strconv.ParseFloat(numFloat, 64)
	fmt.Println(num4)
	fmt.Printf("Data type is %T\n", num4)

	//
	// -----------------------------------------------------------------
	//
	// Call from arraydata
	//
	arraydata.Array()

	//
	// call slice
	//
	slice.SliceCustom()
	callSlice := make([]int, 20)

	// print type of callslice
	fmt.Printf("%T\n", callSlice)

	// call function with value and print return value
	fmt.Println(slice.SliceModifier(callSlice))

}
