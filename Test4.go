package main

import (
	"fmt"
)

/**
测试数组
*/
func main() {
	//方式1
	var nums [3]int
	nums[0] = 0
	nums[1] = 1
	fmt.Println("方式1：")
	fmt.Println(nums)

	//方式2
	var nums1 = []int{0, 1, 2, 3, 4}
	var index, num int
	for i := 0; i < len(nums1); i++ {
		nums1[i] = i * i
	}
	fmt.Println("方式2：")
	//遍历
	for index, num = range nums1 {
		fmt.Println(index, num)
	}

	//方式3
	var nums2 = [5]int{0, 1, 2, 3, 4}
	fmt.Println("方式3：")
	for i := 0; i < len(nums2); i++ {
		fmt.Println(i, nums2[i])
	}

	//打印杨辉三角
	var arr []int //初始行
	for i := 0; i < 10; i++ {
		arr = getYHSJ(arr)
		fmt.Println(arr)
	}

}

//杨辉三角 返回每一行得数
func getYHSJ(arr []int) []int {
	var res []int
	arrLen := len(arr)
	//以1开始
	res = append(res, 1)
	if 0 == arrLen {
		//遇到第一行
		return res
	}

	//使用上一行数据计算本行数据
	for i := 0; i < arrLen-1; i++ {
		res = append(res, arr[i]+arr[i+1])
	}
	//以1结束
	res = append(res, 1)
	return res
}
