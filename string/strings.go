package main

import (
	"fmt"
	"strings"
)

func main1() {
	str := "hello world                    !!!"
	/*
		//EqualFold  判断两个utf-8编码字符串（将unicode大写、小写、标题三种格式字符视为相同）是否相同。忽略大小写
		fmt.Println("字符串相等比较EqualFold",strings.EqualFold(str,"helloworld!!!"),strings.EqualFold(str,"Helloworld!!!"))
		//判断s是否有前缀字符串prefix 和 suffix。
		fmt.Println("前缀判断HasPrefix",strings.HasPrefix(str,"hello"))
		fmt.Println("后缀判断HasSuffix",strings.HasSuffix(str,"!!!"))
		//判断字符串s是否包含子串substr。
		fmt.Println("是否含有指定字符串Contains",strings.Contains(str,"llo"))
		//判断字符串s是否包含utf-8码值r  r为字符
		fmt.Println("是否含有指定字符ContainsRune",strings.ContainsRune(str,'o'))
		fmt.Println("指定字符串中是否有包含的字符ContainsAny",strings.ContainsAny(str,"lmn"))
		//Count 统计重复数量
		fmt.Println("Count",strings.Count(str,"!"),strings.Count(str,"!!"))
		fmt.Println("Index",strings.Index(str,"e"))
		fmt.Println("Index",strings.Index(str,"e"))
	*/
	fmt.Println("字符c在s中第一次出现的位置  IndexByte", strings.IndexByte(str, 'l'))
	fmt.Println("unicode码值r在s中第一次出现的位置  IndexRune", strings.IndexRune(str, 'l'))
	fmt.Println("unicode码值r在s中第一次出现的位置  IndexRune", strings.IndexAny(str, "dfdsalo"))
	//fmt.Println("s中第一个满足函数f的位置i（该处的utf-8码值r满足f(r)==true），不存在则返回-1。",strings.IndexFunc(str,true))
	//fmt.Println("使用_case规定的字符映射，返回将所有字母都转为对应的小写版本的拷贝",strings.ToLowerSpecial(str,true))
	//fmt.Println("使用_case规定的字符映射，返回将所有字母都转为对应的大写版本的拷贝。",strings.ToUpperSpecial(str,true))
	//fmt.Println("使用_case规定的字符映射，返回将所有字母都转为对应的标题版本的拷贝。",strings.ToTitleSpecial(str,true))
	//fmt.Println("将s的每一个unicode码值r都替换为mapping(r)，返回这些新码值组成的字符串拷贝。如果mapping返回一个负值，将会丢弃该码值而不会被替换。（返回值中对应位置将没有码值）",strings.Map(str,true))
	//TrimFunc  TrimLeftFunc TrimRightFunc  FieldsFunc
	//LastIndex  LastIndexAny  LastIndexFunc  ToLower  Replace  TrimLeft  TrimRight  TrimSuffix  Join
	fmt.Println("返回s中每个单词的首字母都改为标题格式的字符串拷贝", strings.Title(str))
	fmt.Println("返回s中每个单词的首字母都改为标题格式的字符串拷贝", strings.ToTitle(str))
	fmt.Println("返回count个s串联的字符串。", strings.Repeat(str, 9))
	fmt.Println("返回将s前后端所有cutset包含的utf-8码值都去掉的字符串。", strings.Trim(str, "!"))
	fmt.Println("返回将s前后端所有空白（unicode.IsSpace指定）都去掉的字符串。", strings.TrimSpace(str))
	fmt.Println("返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串。如果字符串全部是空白或者是空字符串的话，会返回空切片。", strings.Fields(str))
	fmt.Println("用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。", strings.Split(str, ""))
	fmt.Println("用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。参数n决定返回的切片的数目：", strings.SplitAfterN(str, "d", 2))
	//  SplitAfter  SplitAfterN

	/*
		Replacer  type
		Reader	  type  ？？？
	*/
}
