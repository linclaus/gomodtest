package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Barrier struct {
	currentCnt         int
	threshold          int
	finishedSignalChan chan struct{}
	lock               sync.Mutex
}

func NewBarrier(threshold int) *Barrier {
	return &Barrier{
		currentCnt:         0,
		threshold:          threshold,
		finishedSignalChan: make(chan struct{}),
	}
}

func (b *Barrier) addOne() (barrierNum int) {
	b.lock.Lock()
	barrierNum = b.currentCnt
	fmt.Printf("Barrier:%d added\n", barrierNum)
	b.currentCnt++
	b.lock.Unlock()
	return
}

func (b *Barrier) BarrierWait() {
	barrierNum := b.addOne()
	if b.currentCnt == b.threshold {
		close(b.finishedSignalChan)
	} else {
		select {
		case <-b.finishedSignalChan:
			break
		}
	}
	fmt.Printf("Barrier:%d finished\n", barrierNum)
}

func TestMain(t *testing.T) {
	var wg sync.WaitGroup
	br := NewBarrier(10)
	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			br.BarrierWait()
		}()
	}
	wg.Wait()
	time.Sleep(1000)
	fmt.Println("Main ended")
}
