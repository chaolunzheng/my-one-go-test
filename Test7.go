package main

import "fmt"

/**
测试go中的字典map，以及其迭代方式
 */
func main(){
	//key不能重复
	//Map 是一种无序的键值对的集合。无序的
	// Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值
	var dict= map[string]string{"a":"1","b":"2"}
	//增加元素
	dict["c"]="3"
	for key,val:=range dict{
		fmt.Printf("key:%s->val:%s\n",key,val)
	}

	//空map
	dict2:=make(map[string]int)
	fmt.Println(dict2)
	dict2["美国"]=10
	dict2["中国"]=100
	dict2["韩国"]=0
	dict2["日本"]=-1
	for key,val:=range dict2{
		fmt.Printf("key:%s->val:%d\n",key,val)
	}
	//查看是否存在，并返回值
	val,is:= dict2["中国"]
	if is {
		fmt.Println(val)
	}else{
		fmt.Println("no")
	}

	//delete函数使用
	delete(dict2,"日本")
	for key,val:=range dict2{
		fmt.Printf("key:%s->val:%d\n",key,val)
	}
}

