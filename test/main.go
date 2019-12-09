package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println("go 我来罗")
	//bool 值
	var flag1 bool = true
	//整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码
	//int  int8-64  uint8-64
	var i8 int8 = -128
	var i16 int16 = 9
	var i32 int32 = 9
	var i64 int64 = 9
	fmt.Println(i8)
	fmt.Println(i16)
	fmt.Println(i32)
	fmt.Println(i64)
	//float float32-64  complex32-64
	fmt.Sprintln(flag1)
	const name = 0
	fmt.Println(name)
	var pig person
	pig.name, pig.age = "pig", 8
	monkey := person{name: "monkey", age: 7}
	fmt.Println("monkey的姓名,%s", monkey.age)
	old, young := older(pig, monkey)
	fmt.Printf("old---%s,young----%s", old.name, young.name)
	zhangsan := person{animal{"yellow"}, "zhangsan", 99}
	fmt.Println(zhangsan.animal.color)

	/**
	标识符：数字 字幕  下划线  首字符字幕和下划线开头
	*/

}

type animal struct {
	color string
}
type person struct {
	animal
	name string
	age  int
}

func older(p1, p2 person) (person, person) {
	if p1.age > p2.age {
		return p1, p2
	}
	return p2, p1
}
