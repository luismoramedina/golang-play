package main

import "fmt"
import "math"
import "strconv"

func main() {
	x := 3
	s := []string{"a", "b"}
	m := map[string]string {}
	fmt.Println(x)
	fmt.Println(s)
	m["a"] = "b"
	m["c"] = "d"
	fmt.Println(m)
	fmt.Println(math.Pi)
	fmt.Println(len(m))
	for value := range m {
		fmt.Println(value)
	}

	st := "1"
	value, _ := strconv.Atoi(st)
	fmt.Println("value", value)

	for position, value := range s {
		fmt.Println(position, value)
	}



}
