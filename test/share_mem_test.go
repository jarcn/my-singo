package test

import (
	"encoding/base64"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestConunter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++ //线程不安全
		}()
	}
	time.Sleep(time.Second * 1)
	fmt.Println(counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock() //加锁互斥使用共享内存
			counter++
		}()
	}
	time.Sleep(time.Second * 1)
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait() //等待所有的携程执行完
	t.Logf("counter = %d", counter)

}

func TestBase64(t *testing.T) {
	data := "SB-Mid-server-qgGG6qKkEwylnvYhE44WddNQ"
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
}

func TestGorutine(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second * 1)
}
