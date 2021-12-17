package test

import (
	"fmt"
	"sync"
	"testing"
)

func dataP(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataR(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestPR(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataP(ch, &wg)
	wg.Add(1)
	dataR(ch, &wg)
	wg.Add(1)
	dataR(ch, &wg)
	wg.Wait()
}
