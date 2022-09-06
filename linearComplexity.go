package main

import "fmt"

func main() {
	var m [10]int

	for i := 0; i < 10; i++ {
		m[i] = i * 200
		fmt.Printf("Element[%d] = %d\n", i, m[i])
	}

}
