//
// always use before function call statement
// defer keyword execute last of the function call
// multiple defer works like LIFO - Stack
//

package main

import "fmt"

func main() {
	fmt.Println("Start")
	defer fmt.Println("Middle")
	defer fmt.Println("Middle executes first")
	fmt.Println("Last")

}
