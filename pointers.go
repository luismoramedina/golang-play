package main

import "fmt"

func main() {
	x := 5
	pointers(&x)
	fmt.Print(x)
}

func pointers(x *int)  {
	*x = 0
}


