package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{3, 1, 6}
	sort.Ints(slice)
	fmt.Println(slice)
	slice1 := []string{"3", "8", "1"}
	sort.Strings(slice1)
	fmt.Println(slice1)
	fmt.Println(len(slice))
}
