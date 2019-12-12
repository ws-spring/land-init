package main

import (
	"fmt"
	"runtime"
	"sync"
)

func sum1(num int,ch chan int)(int){
	total:=0;
	for i:=0;i<num;i++{
		total+=i;
	}
	ch <- total
	return total
}

func sum2(num int)(int){
	total:=0;
	for i:=0;i<num;i++{
		total+=i;
	}
	return total
}


func  write(num int,ch chan int) {
	for i:=0;i<num;i++{
		ch <- i
	}
}

func  read(ch chan int) {
	for{
		var b int
		b = <-ch
		fmt.Println(b)
	}
}


func main() {
	// intChan := make(chan int, 10)
	// go sum1(10,intChan);
	// res:= <- intChan
	// fmt.Println(res)
	// fmt.Println(sum2(10))
	intChan2 := make(chan int, 10)
	//在main方法执行routine 不一定执行
	fmt.Println("cou核心数cpus:", runtime.NumCPU())
    fmt.Println("环境位置goroot:", runtime.GOROOT())
    fmt.Println("操作系统archive:", runtime.GOOS)
    fmt.Println("退出当前:", runtime.Goexit)
    fmt.Println("设置最大的可同时使用的 CPU 核数:", runtime.GOMAXPROCS)
    
	go write(5,intChan2 )
	fmt.Println("返回正在执行和排队的任务总数:", runtime.NumGoroutine)
	
	runtime.Gosched()
	fmt.Println("99999999999999999999")
	go read(intChan2)
	sync.WaitGroup()
	//time.Sleep(10 * time.Second)

	
}
