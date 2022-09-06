//elements are added through the PushBack method on the list,
//which is in the container/list package

package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(35)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}

}
//The list is iterated through the for loop, 
//and the element's value is accessed through the Value method.
