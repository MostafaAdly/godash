package main

import (
	"fmt"
	"main/src/Object/List"
)

func main() {
	arr :=
		List.NewArray([]int{1, 2, 3, 4, 5}).Map(func(value int, index int) int {
			return value * 2
		}).Build()
	fmt.Println(arr)
}
