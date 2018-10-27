package main

import (
	"errors"
	"fmt"
	"math"
)

/**
异常处理
 */

func main() {
	//if的写法，很厉害了
	if res, err := Sqrt(-9); err == nil {
		fmt.Printf("res:%f,没有错误", res)
	} else {
		fmt.Printf("有错误err:%v", err)
	}
}

//函数
func Sqrt(num float64) (float32, error) {
	errs := `
		错误啦
		错误啦
		错误啦
		错误啦
		错误啦
	`
	if num < 0 {
		return -1, errors.New(errs)
	}
	return float32(math.Sqrt(num)), nil
}
