package test

import (
	"errors"
	"testing"
)

// 1,1,2,3,5,8
func TestFibonacci(t *testing.T) {
	fan := []int{1, 1}
	for i := 2; i < 8; i++ {
		fan = append(fan, fan[i-2]+fan[i-1])
	}
	t.Log(fan)
}

func genFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, errors.New("n should not be less than 2")
	}
	fib := []int{1, 1}
	for i := 2; i < 8; i++ {
		fib = append(fib, fib[i-2]+fib[i-1])
	}
	return fib, nil
}

// go 中的 try_catch
// 使用 recover 方式恢复异常 是否是一个好的编程习惯有待商榷
// 例如 系统 health check 会因为异常恢复机制导致僵尸服务
func TestGen(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("recover from", err)
		}
	}()
	if v, err := genFibonacci(1); err != nil {
		panic(err)
	} else {
		t.Log(v)
	}
}
