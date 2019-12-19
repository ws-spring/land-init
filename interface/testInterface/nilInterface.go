package main

    import (
        "fmt"
    )
	//空接口的默认值  nil

    type People interface {  
        GetName() string
    }

    // 输出 "cbs is nil 类型"
    func main() {  
        var cbs People
        if cbs == nil {
            fmt.Println("cbs is nil 类型")  
        }
    }
