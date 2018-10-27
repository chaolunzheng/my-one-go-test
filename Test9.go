package main

import (
	"fmt"
)

/*
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，
任何其他类型只要实现了这些方法就是实现了这个接口
 */

 //接口
type FlyInf interface {
	say()
	fly()
}
//结构体
type BirdBlack struct{
	sing string
	flyHeight int
}

type BirdWhite struct{
	sing string
	flyHeight int
}

//实现方法
func (bird BirdBlack) fly(){
	fmt.Println("BirdBlack能飞100米高",bird.flyHeight)
}

func (bird BirdBlack) say(){
	fmt.Println("BirdBlack我😁哈哈的叫。。。",bird.sing)
}

func (bird BirdWhite) fly(){
	fmt.Println("BirdWhite能飞111米高")
}

func (bird BirdWhite) say(){
	fmt.Println("BirdWhite我嘎嘎的叫。。。")
}

func main() {
	//实现多态
	var bird0= new(BirdBlack)
	bird0.sing="hhhhhhhhhhhhhh"
	bird0.flyHeight=10000
	var bird1 = new(BirdWhite)
	bird1.sing="xixi"

	bird0.say()
	bird1.say()
	bird0.fly()
	bird1.fly()

}
