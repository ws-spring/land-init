package main

import (
	"fmt"
)

//https://www.cnblogs.com/liuzhongchao/p/9159896.html
func main() {
	//数组  数组是值类型的
	arr1 := [...]int{2, 5, 7, 8, 9, 9, 0}
	//arr2:=[6]int{3}
	//arr3:=[6]int{}
	//fmt.Println(len(arr1),arr1, cap(arr3))
	//for i,v:=range arr2{
	//	fmt.Println(i,v)
	//}
	//for i,v:=range(arr2){
	//	fmt.Println(i,v)
	//}
	//for i:=0;i<len(arr1);i++{
	//	fmt.Println(i,arr1[i])
	//}
	//切片slice  相当于list
	//len 为切片的长度    cap  为切片的容量
	slice1 := arr1[:]
	slice2 := make([]int, 3, 4)
	fmt.Println(slice1)
	for i, _ := range slice1 {
		slice1[i] = slice1[i] * 2
	}
	fmt.Println(slice1)
	fmt.Println(slice2)

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Println("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6

	fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))

	//切片长度变化规律：切片的容量不是一直以2倍数进行扩容的  扩大到*512后，扩大的规则就变了
	var cars = []string{"BWM", "benchi", "jetta"}
	slice3 := cars
	fmt.Println("原始值", slice3, len(slice3), cap(slice3))
	//for i:=0;i<1000000;i++ {
	//	cappacity:=cap(slice3)
	//	for j := 0; j < 3; j++ {
	//		slice3 = append(slice3, strconv.Itoa(i))
	//	}
	//	if(cappacity!=cap(slice3)){
	//		fmt.Println("增加值",len(slice3), cap(slice3))
	//	}
	//}
	//空值长度为0，为空值长度也为0
	var slice4 []string
	slice5 := []string{}
	if slice4 == nil {
		fmt.Println("切片的长度----", slice4, len(slice4))
		fmt.Println("改切片是空的")
	}
	if slice5 != nil {
		fmt.Println("切片的长度----", slice4, len(slice4))
		fmt.Println("该切片不是空的")
	}

	//切片作为函数参数

}

//切片作为函数参数
type slice struct {
	Length        int
	Capacity      int
	ZerothElement *byte //指向首元素的指针
}
