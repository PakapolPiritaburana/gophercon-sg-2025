package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// สร้าง semaphore ที่อนุญาตให้ทำงานพร้อมกัน 3 งาน
func basicSemaphore() {
	semaphore := make(chan struct{}, 3) // ขนาด buffer = จำนวนที่อนุญาต

	var wg sync.WaitGroup

	// สร้าง 10 งาน แต่จะทำงานพร้อมกันได้แค่ 3 งาน
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Acquire semaphore
			semaphore <- struct{}{}
			defer func() { <-semaphore }() // Release semaphore

			// Critical section - ทำงานที่ต้องจำกัดจำนวน
			fmt.Printf("Worker %d กำลังทำงาน\n", id)
			time.Sleep(2 * time.Second)
			fmt.Printf("Worker %d เสร็จแล้ว\n", id)
		}(i)
	}

	wg.Wait()
}

func semaphoreWithTimeout() {
	semaphore := make(chan struct{}, 2)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// พยายาม acquire semaphore พร้อม timeout
			select {
			case semaphore <- struct{}{}:
				defer func() { <-semaphore }()

				// ทำงาน
				fmt.Printf("Worker %d เริ่มทำงาน\n", id)
				time.Sleep(3 * time.Second)
				fmt.Printf("Worker %d เสร็จ\n", id)

			case <-ctx.Done():
				fmt.Printf("Worker %d timeout\n", id)
				return
			}
		}(i)
	}

	wg.Wait()
}

func main() {
	//basicSemaphore()
	semaphoreWithTimeout()
}
