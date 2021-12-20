package test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 5)
	return fmt.Sprintf("the result is from %d", id)
}

func FirstResp() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(id int) {
			ret := runTask(id)
			ch <- ret
		}(i)
	}
	return <-ch
}

func AllResp() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(id int) {
			ret := runTask(id)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for i := 0; i < numOfRunner; i++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestFirstResp(t *testing.T) {
	t.Log("before:", runtime.NumGoroutine())
	// t.Log(FirstResp())
	t.Log(AllResp())
	time.Sleep(time.Second * 1)
	t.Log("after:", runtime.NumGoroutine())
}
