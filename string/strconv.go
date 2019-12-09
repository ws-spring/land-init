package main

import (
	"fmt"
	"strconv"
)

/**
func  多个返回值  需要理解
*/
func main() {
	//str:="helloworl"
	/**
	Constants  Variables  NumError  func (*NumError) Error
	*/
	fmt.Println("返回一个字符是否是可打印的，和unicode.IsPrint一样，r必须是：字母（广义）、数字、标点、符号、ASCII空格。", strconv.IsPrint('国'))
	fmt.Println("返回字符串s是否可以不被修改的表示为一个单行的、没有空格和tab之外控制字符的反引号字符串。", strconv.CanBackquote("我爱你\r\n中国"))
	fmt.Println("返回字符串s在go语法下的双引号字面值表示，控制字符、不可打印字符会进行转义。（如\t，\n，\xFF，\u0100）", strconv.Quote("我爱你\r\n中国"))
	fmt.Println("返回字符串s在go语法下的双引号字面值表示，控制字符和不可打印字符、非ASCII字符会进行转义。", strconv.QuoteToASCII("我爱你\r\n中国"))
	fmt.Println("返回字符r在go语法下的单引号字面值表示，控制字符、不可打印字符会进行转义。（如\t，\n，\xFF，\u0100）", strconv.QuoteRune('国'))
	fmt.Println("返回字符r在go语法下的单引号字面值表示，控制字符、不可打印字符、非ASCII字符会进行转义。", strconv.QuoteRuneToASCII('国'))
	//func Unquote(s string) (t string, err error)
	/**
	func Unquote(s string) (t string, err error)
	函数假设s是一个单引号、双引号、反引号包围的go语法字符串，解析它并返回它表示的值。（如果是单引号括起来的，函数会认为s是go字符字面值，返回一个单字符的字符串）
	*/
	/**
	func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
	函数假设s是一个表示字符的go语法字符串，解析它并返回四个值：

	1) value，表示一个rune值或者一个byte值
	2) multibyte，表示value是否是一个多字节的utf-8字符
	3) tail，表示字符串剩余的部分
	4) err，表示可能存在的语法错误
	quote参数为单引号时，函数认为单引号是语法字符，不接受未转义的单引号；双引号时，函数认为双引号是语法字符，不接受未转义的双引号；如果是零值，函数把单引号和双引号当成普通字符。
	*/
	/*
		func ParseBool(str string) (value bool, err error)
		返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
	*/

}
