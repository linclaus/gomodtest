package main

import (
	"fmt"
	"sync"
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
	defer b.lock.Unlock()
	barrierNum = b.currentCnt
	fmt.Printf("Barrier:%d added\n", barrierNum)
	b.currentCnt++
	return
}

func (b *Barrier) BarrierWait() {
	barrierNum := b.addOne()
	if b.currentCnt == b.threshold {
		// Broadcast the ready signal
		close(b.finishedSignalChan)
	} else {
		// Blocked until the ready signal
		<-b.finishedSignalChan
	}
	fmt.Printf("Barrier:%d finished\n", barrierNum)
}

func main() {
	var wg sync.WaitGroup
	br := NewBarrier(5)
	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			br.BarrierWait()
		}()
	}
	wg.Wait()
	time.Sleep(1 * time.Second)
	fmt.Println("Main ended")
}
