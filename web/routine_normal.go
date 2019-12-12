//https://www.cnblogs.com/chenny7/p/4498322.html

package  main

import "fmt"

func Add(x, y int) {
    z := x + y
    fmt.Println(z)
}

func main() {
    for i:=0; i<10; i++ {
        go Add(i, i)
    }
}