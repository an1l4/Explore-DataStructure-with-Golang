//power series of integers are calculated and the square and cube of the integer is returned as a tuple

package main

import "fmt"

func powerSeries(a int) (int, int, error) {
	return a * a, a * a * a, nil

}

func main() {
	square, cube, error := powerSeries(3)
	fmt.Println("Square: ", square, "Cube: ", cube)
	fmt.Println(error)

}
