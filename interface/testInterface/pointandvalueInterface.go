    package main
    /**
    指针与值类型实现接口的区别
    */

    import (
        "fmt"
    )

    type People interface {
        ReturnName() string
    }

    type Student struct {
        Name string
    }

    type Teacher struct {
        Name string
    }

    func (s Student) ReturnName() string {
        return s.Name
    }

    func (t *Teacher) ReturnName() string {
        return t.Name
    }

    func main() {
        cbs := Student{Name: "咖啡色的羊驼"}
        sss := Teacher{Name: "咖啡色的羊驼的老师"}

        // 值类型
        var a People
        a = cbs 
        name := a.ReturnName()
        fmt.Println(name)

        // 指针类型
        // a = sss <- 这样写不行！！！
        a = &sss // 由于是指针类型，所以赋值的时候需要加上&
        name = a.ReturnName()
        fmt.Println(name) // 输出"咖啡色的羊驼的老师"
    }