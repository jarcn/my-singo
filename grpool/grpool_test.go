package grpool

import (
	"runtime"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	println("using MAXPROC")
	numCpus := runtime.NumCPU()
	runtime.GOMAXPROCS(numCpus)
}

func TestNewWorker(t *testing.T) {
	pool := make(chan *worker)
	worker := newWorker(pool)
	worker.start()
	assert.NotNil(t, worker)
	worker = <-pool
	assert.NotNil(t, worker, "worker should register itself to the pool")
	called := false
	done := make(chan bool)
	job := func() {
		called = true
		done <- true
	}
	worker.jobChannel <- job
	<-done
	assert.Equal(t, true, called)
}

func TestNewPool(t *testing.T) {
	pool := NewPool(1000, 10000)
	defer pool.Release()

	iterations := 1000000
	pool.WaitCount(iterations)
	var counter uint64 = 0
	for i := 0; i < iterations; i++ {
		arg := uint64(1)
		job := func() {
			defer pool.JobDone()
			atomic.AddUint64(&counter, arg)
			assert.Equal(t, uint64(1), arg)
		}
		pool.JobQueue <- job
	}
	pool.WaitAll()
	counterFinal := atomic.LoadUint64(&counter)
	assert.Equal(t, uint64(iterations), counterFinal)
}
