package e

import (
	"errors"
	"testing"
)

var LessThanTwoError = errors.New("小于2了")
var LargerThanHundredError = errors.New("大于100了")

func GetFibonacci(n int) ([]int, error) {

	if n < 2 {
		return nil, LessThanTwoError
	}

	if n > 100 {
		return nil, LargerThanHundredError
	}

	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}

	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(101); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}
