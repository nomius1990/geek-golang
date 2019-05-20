package series

import (
	"fmt"
)

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

//大写的声明在包外可以访问
func GetFibonacciSerie(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

//小写的声明在包外无法访问
func square(n int) int {
	return n * n
}
