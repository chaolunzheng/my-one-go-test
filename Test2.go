package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//1 if条件语句
	cond := 50
	if cond < 40 {
		fmt.Printf("条件是cond=%d 成立\n", cond)
	} else {
		fmt.Printf("条件是cond=%d 不成立\n", cond)
	}

	//2 switch语句
	switch cond {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}

	switch {
	case true:
		fmt.Println("true")
	}

	//3. select
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := <-c3: // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

	//4 TODO 死循环
	for {
		fmt.Println(1)
		if 1 == 1 {
			break
		}
	}

	//while 直到条件满足
	i := 0
	for i <= 10 {
		fmt.Print(i)
		i++
	}
	fmt.Println()

	//for循环
	for i := 0; i < 3; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	//获取命令行参数1
	fmt.Println("method1:", strings.Join(os.Args[1:], " "))

	//获取命令行参数2
	var s, sep string
	for j := 1; j < len(os.Args); j++ {
		s += sep + os.Args[j]
		sep = " "
	}
	fmt.Println("method2:", s)

	//获取命令行参数3
	var res, sep2 string
	for _, s = range os.Args[1:] {
		res += sep2 + s
		sep2 = " "
	}
	fmt.Println("method3:", res)

	//利用for循环嵌套求素数->大于1的整数中只能被1和它本身整除的数
	fmt.Println("100以内的素数：")
	var m, n int
	for m = 2; m < 100; m++ {
		for n = 2; n <= m/n; n++ {
			if m%n == 0 {
				break
			}
		}
		if n > m/n {
			fmt.Print(" ", m)
		}
	}
	fmt.Println()
	for m := 2; m < 100; m++ {
		for n := 2; n<=m; n++ {
			if m%n == 0 {
				if n<m{
					break
				}else{
					fmt.Print(" ",m)
				}
			}
		}
	}


}
