package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Barrier struct {
	currentCnt         int32
	threshold          int32
	finishedSignalChan chan struct{}
	lock               sync.Mutex
}

func NewBarrier(threshold int32) (*Barrier, error) {
	if threshold < 0 {
		return nil, fmt.Errorf("Create barrier error because of negative threshold: %d", threshold)
	}
	return &Barrier{
		currentCnt:         0,
		threshold:          threshold,
		finishedSignalChan: make(chan struct{}),
	}, nil
}

func (b *Barrier) BarrierWait() {
	barrierNum := atomic.AddInt32(&b.currentCnt, 1)
	if barrierNum > b.threshold {
		// Other cases and barrier will do nothing
		return
	}
	fmt.Printf("Barrier:%d added\n", barrierNum)
	if barrierNum == b.threshold {
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
	threshold := 5
	br, err := NewBarrier(int32(threshold))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	wg.Add(threshold)
	for i := 0; i < threshold; i++ {
		go func() {
			defer wg.Done()
			br.BarrierWait()
		}()
	}
	wg.Wait()
	fmt.Println("Main ended")
}
