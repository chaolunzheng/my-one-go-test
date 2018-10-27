package main

import "fmt"

/**
函数示例
 */


 //闭包1 返回一个函数
func getSeq() func() int{
	i:=0
	return func() int {
		i++
		return i
	}
}

//闭包2  带参数
func addFUnc(x1 int,x2 int) func(x3 int,x4 int) (int,int,int){
	i:=0
	return func(x3 int,x4 int) (int,int,int){
		i++
		return i,x1+x2,x3+x4
	}
}

//结构体类型 圆
type Circle struct {
	//半径
	radius float64
}

//方法演示 求圆的面积 该 method 属于 Circle 类型对象中的方法
func (c Circle) area() float64{
	return 3.14*c.radius*c.radius
}
func main(){
	//函数作为值传递
	add:= func(a int,b int) int{
		return a+b
	}
	fmt.Println(add(3,7))

	//闭包演示
	num1:=getSeq() //闭包1
	fmt.Println(num1())
	fmt.Println(num1())
	fmt.Println(num1())

	num2:=getSeq() //闭包2
	fmt.Println(num2())
	fmt.Println(num2())

	//闭包2 带参数
	add2:=addFUnc(8,3)
	fmt.Println(add2(1,9))
	fmt.Println(add2(6,4))


	//方法演示
	var c Circle
	c.radius=10.0
	areas := c.area()
	fmt.Println("圆的面积",areas)

}