package main

import "fmt"

var(
	as =new([3]int)
)

func main() {
	/**
	冒泡排序
	 */
	arr:= [10]int{4, 4, 1, 2, 12, 5, 6, 834, 3, 0}

	//var arr = new([10]int)
	fmt.Println(arr)

	for i := len(arr)-1; i >=0; i-- {
		for j := 0; j < i; j++ {
			if arr[j]>arr[j+1]{
				tem := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tem
			}
		}
	}
	fmt.Println(&arr)

	var ss Stu = Stu{"string"}

	ss.print()
}



type Stu struct {
	text string
}
//给Stu类型绑定方法
func (s Stu) print(){
	text := s.text
	fmt.Println("text:",text)
}
