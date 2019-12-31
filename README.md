# land-init


1.make和new的区别  https://www.cnblogs.com/vincenshen/p/9356974.html
相同点：分配内存
不同点：
new： 被用来分配除了引用类型的所有其他类型的内存： int, string, array等，作用是分配内存地址和初始化内存地址指向的值  func new(Type) *Type，返回的是指向类型的指针。
make：给引用类型分配内存，如channel，map和切片，返回的类型就是这三个类型本身，而不是他们的指针类型   func make(t Type, size ...IntegerType) Type，返回的是这三个引用类型本身；

2.&与*的区别：
& 是取地址符号 , 即取得某个变量的地址 , 如 ; &a
* 是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值 .

3.结构体的实例化
a.创建指针类型的结构体：ins := new(T)
	其中：
	T 为类型，可以是结构体、整型、字符串等。
	ins：T 类型被实例化后保存到 ins 变量中，ins 的类型为 *T，属于指针。

b.取结构体的地址实例化：ins := &T{}
	其中：
	T 表示结构体类型。
	ins 为结构体的实例，类型为 *T，是指针类型。
c.基本实例化：var ins T

4.time

Golang神奇的2006-01-02 15:04:05:https://www.jianshu.com/p/c7f7fbb16932

时间可分为时间点与时间段，golang 也不例外，提供了以下两种基础类型
- 时间点(Time)
- 时间段(Duration)

除此之外也提供了以下类型，做一些特定的业务
- 时区(Location)
- Ticker
- Timer(定时器)


4.net/http   https://blog.csdn.net/linkvivi/article/details/80250602
步骤：
1.定义处理器： 
	func initHandler(w http.ResponseWriter,r *http.Request)
	
2.在路由器中注册处理器  
	http.HandleFunc("/", initHandler)
	
3.实例化serve对象(设置监听地址和端口)：
	err := http.ListenAndServe(":9090", nil) 
	
	server := &Server{Addr: addr, Handler: handler}
    server.ListenAndServe()
	源码：
	func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
		DefaultServeMux.HandleFunc(pattern, handler)
	}

疑问：
defer
panic
sync.RWMutex






