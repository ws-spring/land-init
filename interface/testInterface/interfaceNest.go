    package main
	/**
	9.接口嵌套
	类似于PHP的接口继承，Golang也有它的接口嵌套。
	*/

    import (
        "fmt"
    )

    type People interface {
        ReturnName() string
    }

    type Role interface {
        People // 接口嵌套
        ReturnRole() string
    }

    type Student struct {
        Name string
    }

    func (s Student) ReturnName() string {
        return s.Name
    }

    func (s Student) ReturnRole() string {
        return "学生"
    }



    func main() {
        cbs := Student{Name: "咖啡色的羊驼"}

        var a Role
        a = cbs 
        name := a.ReturnName()
        fmt.Println(name)

        role := a.ReturnRole()
        fmt.Println(role) 
    }