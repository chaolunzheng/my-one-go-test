package main

import (
	"awesomeProject/test"
	"fmt"
	"reflect"
)

/**
golang的反射讲解
 */
type Student struct {
	/**
	学生类型
	 */
	Name  string
	Age   int
	Sex   int
	Email string
}
//接口定义
type Communication interface {
	Say(hello string)
	ByBy(byby string)
}

//学生的方法 实现了接口Communication（需要实现接口所有方法 就实现了接口 不用显示声明实现了某接口
func (s Student) Say(hello string){
	fmt.Println(hello)
}

func (s Student) ByBy(byby string){
	fmt.Println(byby)
}

//绑定方法
func (s Student) PlayPhone(playPhone string){
	fmt.Println(playPhone)
}



func main() {
	s := Student{Name: "雪穗", Age: 25, Sex: 0, Email: "alun@163.com"}

	fmt.Println(s)
	c:=s//子类指向父类的引用
	c.ByBy("byby")
	c.Say("say hello")
	c.PlayPhone("playPhone")

	fmt.Println("--------------反射调用demo----field----------")

	sv := reflect.ValueOf(s)

	//根据属性名字获取属性值
	for i := 0; i < sv.NumField(); i++ {
		f0 := sv.Field(i)
		if !f0.IsValid(){
			//判断是否有该属性
			fmt.Println("IsValid")
			return
		}
		fmt.Println(f0)
	}

	fmt.Println("--------------反射调用demo----method----------")
	sm := sv.MethodByName("PlayPhone")
	fmt.Println(sm)

	args := []reflect.Value{reflect.ValueOf("&(^*%*^$*@#$火星文")}
	sm.Call(args)


	test.Test()
}
