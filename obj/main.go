package main

import "fmt"

/**
  strut  面向对象的实现
*/

type mathRange struct {
	width, height float64
}

func area(r mathRange) float64 {
	return r.width * r.height
}

func main1() {
	c1 := mathRange{2, 2}
	res1 := area(c1)
	fmt.Println(res1)

}
