    package main

    import (
        "fmt"
    )

    // 定义一个接口
    type People interface {
        ReturnName() string
    }

    // 定义一个结构体
    type Student struct {
        Name string
    }

    // 定义结构体的一个方法。
    // 突然发现这个方法同接口People的所有方法(就一个)，此时可直接认为结构体Student实现了接口People
    func (s Student) ReturnName() string {
        return s.Name
    }

    // 如何测试是否已实现该接口？
    // 由于x.(T)只能是接口类型判断，所以传参时候，传入的是接口类型
    // 为何test的类型可以是一个空接口？埋伏笔下文便知。
    func CheckPeople(test interface{}) {
        if _, ok := test.(People); ok {
            fmt.Printf("Student implements People")
        }
    }


    func main() {
        cbs := Student{Name:"咖啡色的羊驼"}
        CheckPeople(cbs) // Student implements People
    }

    // func main() {
    //     cbs := Student{Name:"咖啡色的羊驼"}

    //     var a People
    //     // 因为Students实现了接口所以直接赋值没问题
    //     // 如果没实现会报错：cannot use cbs (type Student) as type People in assignment:Student does not implement People (missing ReturnName method)
    //     a = cbs       
    //     name := a.ReturnName() 
    //     fmt.Println(name) // 输出"咖啡色的羊驼"
    // }