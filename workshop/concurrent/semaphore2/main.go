package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	semaCh chan struct{}
}

func NewSemaphore(maxConcurrency int) *Semaphore {
	return &Semaphore{
		semaCh: make(chan struct{}, maxConcurrency),
	}
}

func (s *Semaphore) Acquire() {
	s.semaCh <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.semaCh
}

func (s *Semaphore) AcquireWithTimeout(timeout time.Duration) error {
	select {
	case s.semaCh <- struct{}{}:
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("semaphore acquire timeout")
	}
}

// ตัวอย่างการใช้งาน
func useSemaphoreStruct() {
	sem := NewSemaphore(3)
	var wg sync.WaitGroup

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Acquire
			if err := sem.AcquireWithTimeout(2 * time.Second); err != nil {
				fmt.Printf("Worker %d: %v\n", id, err)
				return
			}
			defer sem.Release()

			// ทำงาน
			fmt.Printf("Worker %d ทำงาน\n", id)
			time.Sleep(1 * time.Second)
		}(i)
	}

	wg.Wait()
}

func main() {
	useSemaphoreStruct()
}
