package main

import "fmt"

/**
递归测试
我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中
 */
func main(){
	//斐波那契
	for i:=0;i<=10;i++{
		res:= feibanaqi(i)
		fmt.Printf("第%d个斐波那契数为：%d\n",i,res)
	}

	//阶乘
	for i:=0;i<=10;i++{
		res:= factorial(i)
		fmt.Printf("%d的阶乘为：%d\n",i,res)
	}


	//关于类型转化
	var sum = 432
	var count = 13
	fmt.Println("均值为：",float64(sum)/float64(count))

}
//阶乘
func factorial(i int) int {
	if i<2{return i}
	return i*factorial(i-1)
}
//产生斐波那契数
func feibanaqi(n int) int{
	if n<2{
		return n
	}
	return feibanaqi(n-1)+feibanaqi(n-2)
}


