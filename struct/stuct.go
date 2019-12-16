https://www.cnblogs.com/sparkdev/p/10761825.html
package main

//定义结构体

type Person stuct{
	Name string
	Age int
	Email string
}
//简单的定义   如果一个字段在代码中从来不会被用到，那可以把它命名为 _，即空标识符。
type T stuct{a,b int}

//字段标记  
//在定义结构体时还可以为字段指定一个标记信息
type OtherPerson stuct{
	Name string 'json:"name"'
	Age int 'json:"age"'
	Email string 'json:"email"'
}

func main(){
	//结构体中的字段可以是任何类型，甚至是结构体本身，也可以是函数或者接口。
	//可以声明结构体类型的一个变量，然后像下面这样给它的字段赋值：
	var p Person
	p.Name="nick"
	P.Age=28
	//另外，数组可以看作是一种结构体类型，不过它使用下标而不是具名的字段。
	
	
	//声明结构体类型的变量
	//一旦定义了结构体类型就可以使用这个类型创建值。
	//使用结构类型声明变量，并初始化为零值
	//关键字 var 创建了类型为 Person 且名为 nick 的变量。nick 被称作类型 Person 的一个实例(instance)或对象(object)。注意：当声明变量时，这个变量对应的值总是会被初始化。这个值要么用指定的值初始化，要么用零值(即变量类型的默认值)做初始化。对数值类型来说，零值是 0；对字符串来说，零值是空字符串；对布尔类型，零值是 false。
	var nick Person
	//通过上面的方式声明结构体类型的变量，结构体里的每个字段都会用零值初始化。任何时候，创建一个变量并初始化为零值，习惯上会使用关键字 var。这种用法是为了更明确地表示一个变量被设置为零值。
    //使用结构体字面量声明变量，并初始化为非零值
    //如果希望变量被初始化为某个非零值，可以通过结构体字面量和短变量声明操作符(:=)来创建变量。下面的代码展示了如何声明一个 Persion 类型的变量，并使用某个非零值作为初始值。
	// 声明 Person 类型的变量，并初始化所有字段
	nick := Person{
		Name: "nick",
		Age: 28,
		Email: "nickli@xxx.com",
	}
	//短变量声明操作符(:=)
	//它的作用是声明并且赋值一个变量，其好处是不需要写 var 三个字母，另外不需要写类型，Golang 语言会自动根据赋值的内容确定类型。
	//使用短变量声明操作符也有一些限制，比如不能在函数外面使用，即不能用来声明全局变量。另外短变量声明操作符左边至少得有一个变量是没有定义过的。

	//字面量的两种写法
	
	//结构体字面量可以对结构体类型进行初始化，比如前面介绍过的示例：
	nick := Person{
    Name: "nick",
    Age: 28,
    Email: "nickli@xxx.com",
	}
	//第二种形式没有字段名，只声明对应的值，如下面的代码：
	nick := Person{"nick", 28, "nickli@xxx.com"}
	
	//new函数
	
	//new 是一个用来分配内存的内置函数，但是与 C++ 不一样的是，它并不初始化内存，只是将其置零。也就是说，new(T) 会为类型 T 分配被置零的内存，并且返回它的地址，一个类型为 *T 的值。在 Golang 的术语中，其返回一个指向新分配的类型为 T 的指针，这个指针指向的内容的值为零(zero value)。比如下面的代码：
	nick := new(Person)
	fmt.Printf("%T", nick)
	var nick Person
	fmt.Printf("%T", nick)
	
	//使用 new 函数时，声明变量和分配内存并不需要放在一起，可以先声明一个变量，然后再通过 new 函数为之分配内存，比如下面的写法：
	var nick *Person
	nick = new (Person)

	//new 函数的特点是只能把内存初始化为零值并返回其指针，如果要通过字面量初始化该内存就需要使用混合字面量语法(composite literal syntax)：&T{...}
	//比如下面的写法：
	nick := &Person{
    Name:  "nick",
    Age:   28,
    Email: "nickli@xxx.com",
	}
	//此时 nick 的类型也是 *Person。因此我们可以得出下面的结论：
	//表达式 new(Type) 和 &Type{} 是等价的。
	
	//选择器(selector)
	
	//匿名字段和内嵌结构体
	
	//结构体可以包含一个或多个匿名(或者称为内嵌)字段，即这些字段没有显式的名字。仅指明字段的类型，此时该类型就是字段的名字。匿名字段本身可以是一个结构体类型，即结构体可以包含内嵌的结构体。
	//匿名字段
    //匿名字段和面向对象编程中的继承概念相似，可以被用来模拟类似继承的行为。Golang 中的基础就是通过内嵌或组合来实现的，所以说在 Golang 中组合比继承更受欢迎。比如下面的例子：
	type test struct {
    name string
    age int
    int // 匿名字段
	}
	//内嵌结构体
	//结构体也是一种数据类型，所以它同样可以作为匿名字段使用：
	type Person struct {
    Name  string
    Age     int
    Email string
	}
	type Student struct {
		Person
		StudentID int
	}
	//下面的代码可以声明并初始化 Student 类型的变量：
	st := Student {
    Person:    Person{"jack", 22, "jack@xxx.com"},
    StudentID: 1000,
	}
	//从这个示例可以看出，内层结构体被简单地插入或者内嵌进外层结构体。这种简单的 "继承" 机制使得 Golang 很轻松就能实现从一个或一些类型中继承部分或全部的实现。
}