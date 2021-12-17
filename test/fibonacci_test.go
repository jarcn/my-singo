package test

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
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

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new object")
			return 100
		},
	}
	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	runtime.GC()
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new object.")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func doSomething() string {
	time.Sleep(time.Millisecond * 500)
	return "done"
}

func asyncDo() chan string {
	ret := make(chan string, 1)
	go func() {
		ret <- doSomething()
	}()
	return ret
}

func TestSelectTo(t *testing.T) {
	select {
	case ret := <-asyncDo():
		fmt.Println(ret)
	case <-time.After(time.Millisecond * 100):
		t.Log("time out")
	}

}
