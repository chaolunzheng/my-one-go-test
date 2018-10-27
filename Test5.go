package main

import "fmt"

/**
测试go的指针
 */
func main() {
	var val int32 = 10
	//取地址符
	addr := &val
	fmt.Println(val, "的地址为：", addr)

	//一个指针变量指向了一个值的内存地址。
	//类似于变量和常量，在使用指针前你需要声明指针
	var pointer *int32 = &val
	fmt.Printf("变量%d的地址为:%x\n", val, pointer)
	fmt.Printf("地址指向的值为:%d\n", *pointer)

	/**
	Go 空指针
	当一个指针被定义后没有分配到任何变量时，它的值为 nil。
	 */
	var ptr *string
	var ptr1 *int
	fmt.Println("null pointer:", ptr, ptr1, ptr == nil)

	//测试go的数组指针
	var arrs = []int{1, 2, 3}
	var arrsPtr [3]*int
	for i := 0; i < len(arrs); i++ {
		arrsPtr[i] = &arrs[i]
	}
	for j := 0; j < len(arrs); j++ {
		fmt.Print(j, "_", *arrsPtr[j], "\n")
	}
	fmt.Println()

	//指向指针的指针
	var num =10
	var p1 *int = &num
	//存放指针的指针
	var p2 **int = &p1
	fmt.Printf("这只指向指针的指针%x,值为%x,一重指针的实际值为%d\n",p2,*p2,**p2)

	//传递指针
	var s1,s2=1,10
	fmt.Println(s1,s2)
	swap(&s1,&s2)
	fmt.Println(s1,s2)

	//结构体
	//结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
	book1 :=Books{"go编程入门到精通","高智深","计算机丛书",1}
	book2 :=Books{title:"百年孤独",id:2}
	fmt.Println("book1",book1.title, book1.author, book1.subject, book1.id)

	book2.subject="日本文学"//修改结构体成员属性值
	//访问结构为i成员
	fmt.Println("book2",book2.title, book2.author, book2.subject, book2.id)

	//以结构体作为函数参数传递
	printBook(book1)

	//结构体指针作为函数参数传递
	printBookPtr(&book1)

}
//结构体作为函数参数传递
func printBook(book1 Books) {
	fmt.Println("结构体作为函数参数传递:")
	fmt.Println("book1",book1.title, book1.author, book1.subject, book1.id)
}
//结构体指针作为函数参数传递
func printBookPtr(book1 *Books) {
	fmt.Println("结构体作为函数参数传递:")
	fmt.Println("book1",(*book1).subject, book1.author, book1.subject, book1.id)
}
//交换函数
func swap(a *int, b *int) {
	var temp int
	temp=*a
	*a=*b
	*b=temp
}

//结构体 书
type Books struct{
	//标题
	title string
	//作者
	author string
	//主题分类
	subject string
	//编号
	id int
}
