package test

import (
	"fmt"
	"testing"
	"time"
)

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func service() string {
	time.Sleep(time.Microsecond * 50)
	return "Done"
}

func AsynService() chan string {
	retChan := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retChan <- ret
		fmt.Println("service exited.")
	}()
	return retChan
}

func TestAsynService(t *testing.T) {
	retch := AsynService()
	otherTask()
	fmt.Println(<-retch)
	time.Sleep(time.Second * 1)
}

func TestSyncService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}
