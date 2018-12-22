package main

import "fmt"

/**
测试切片
Go 语言切片是对数组的抽象
*/
func main() {
	//声明一个未指定大小的数组来定义切片
	var slice []int
	//使用make函数初始化切片
	slice = make([]int, 3, 10)
	slice[1] = 1
	fmt.Println(slice)
	//打印出容量
	fmt.Println(cap(slice))

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	//使用数组给切片赋值
	slice = arr[0:]
	fmt.Println(slice)
	//追加元素 动态扩充数组
	ints := append(slice, 6)
	fmt.Println(ints)

	//切片的容量
	fmt.Printf("cap:%d,len:%d\n", cap(slice), len(slice))

	//空切片
	var slice2 []string
	fmt.Println(len(slice2), cap(slice2), slice2 == nil)

	//如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来
	var slice3 []string = make([]string, 2, 3)
	var arrstr = [4]string{"str1", "str2", "str3", "str4"}
	slice3 = arrstr[0:]
	printSlicePro(slice3)

	//扩充切片
	var slice3double = make([]string, len(slice3), cap(slice3)*2)
	/* 拷贝 numbers 的内容到 numbers1 */
	copy(slice3double, slice3)
	printSlicePro(slice3double)
	//获取部分切片
	slice4 := slice3double[1:3]
	printSlicePro(slice4)

	/**
	可以看出切片，实际的是获取数组的某一部分，len切片<=cap切片<=len数组，
	切片由三部分组成：指向底层数组的指针、len、cap。
	*/

	//关于二维数组
	var arr2s = [5][2]int{
		{00, 01},
		{10, 11},
		{20},
		{30},
		{40, 41}}
	fmt.Println(arr2s)
	//遍历二维数组
	for idx, ele := range arr2s {
		fmt.Printf("行索引%d,元素%d\n", idx, ele)
	}
	fmt.Println("------遍历二维数组-------->>")
	for i := 0; i < len(arr2s); i++ {
		for j := 0; j < len(arr2s[i]); j++ {
			fmt.Printf("%d行%d列:%2d\t|\t", i, j, arr2s[i][j])
		}
		fmt.Println()
	}

	/**
	range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
	在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对的 key 值
	*/

}

//打印函数
func printSlicePro(slice []string) {
	fmt.Printf("slice:%v,len:%d,cap:%d\n", slice, len(slice), cap(slice))
}
