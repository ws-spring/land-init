    package main
    /**
    类型判断   if  switch   x.(type)
    */
    import (
        "fmt"
    )
    //既然空接口可以存储任意类型，那么如何区分不同的类型？
    //常用的有两种方法：Comma-ok断言、switch判断。

    // 定义一个结构体
    type Student struct {
        Name string
    }

    // 类型断言
    func main() {
        Params := make([]interface{}, 3)
        Params[0] = 88                   // 整型
        Params[1] = "咖啡色的羊驼"         // 字符串
        Params[2] = Student{Name: "cbs"} // 自定义结构体类型

        // Comma-ok断言
        for index, v := range Params {
            if _, ok := v.(int); ok {
                fmt.Printf("Params[%d] 是int类型 \n", index)
            } else if _, ok := v.(string); ok {
                fmt.Printf("Params[%d] 是字符串类型\n", index)
            } else if _, ok := v.(Student); ok {
                fmt.Printf("Params[%d] 是自定义结构体类型\n", index)
            } else {
                fmt.Printf("list[%d] 未知类型\n", index)
            }
        }

        // switch判断
        for index, v := range Params {
            switch  value := v.(type) {
            case int:
                fmt.Printf("Params[%d] 是int类型, 值：%d \n", index,value)
            case string:
                fmt.Printf("Params[%d] 是字符串类型, 值：%s\n", index,value)
            case Student:
                fmt.Printf("Params[%d] 是Person类型, 值：%s\n", index,value)
            default:
                fmt.Printf("list[%d] 未知类型\n", index)
            } 

        }  
    }