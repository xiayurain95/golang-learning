package code

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/semaphore"
)

type slots struct {
	container                        []int
	consumerPointer, producerPointer int64
	containerMutex                   sync.Mutex
	ctx                              context.Context
	pushMutex                        *semaphore.Weighted
	popMutex                         *semaphore.Weighted
}

func NewSlots(lenth int64) *slots {
	newSlots := &slots{
		container:       make([]int, lenth),
		consumerPointer: lenth - 1,
		producerPointer: lenth - 1,
		ctx:             context.Background(),
		pushMutex:       semaphore.NewWeighted(lenth),
		popMutex:        semaphore.NewWeighted(lenth),
	}
	newSlots.popMutex.Acquire(newSlots.ctx, lenth)
	return newSlots
}

func Comsumer(s *slots) {
	for true {
		s.popMutex.Acquire(s.ctx, 1)
		s.containerMutex.Lock()
		if s.consumerPointer == -1 {
			s.consumerPointer = int64(len(s.container) - 1)
		}
		fmt.Println(s.container[s.consumerPointer])
		s.consumerPointer--
		s.containerMutex.Unlock()
		s.pushMutex.Release(1)
	}
}

func Producer(s *slots) {
	num := 0
	for true {
		s.pushMutex.Acquire(s.ctx, 1)
		s.containerMutex.Lock()
		if s.producerPointer == -1 {
			s.producerPointer = int64(len(s.container) - 1)
		}
		num++
		s.container[s.producerPointer] = num
		s.producerPointer--
		s.containerMutex.Unlock()
		s.popMutex.Release(1)
	}
}
