package main

import (
	"fmt"
	"sync"
)

// stage 1: ส่งเลขเข้าไปใน channel
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// stage 2: รับเลขจาก channel แล้วยกกำลังสองแบบ concurrent
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// stage 3: แสดงผลลัพธ์ที่ได้
func printResults(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}

// fan-in: รวมหลาย channel ให้เหลือ channel เดียว
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// ส่งค่าจาก c ไปยัง out
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// ปิด out channel เมื่อทุก goroutine ทำงานเสร็จ
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// เริ่ม pipeline
	nums := generate(1, 2, 3, 4, 5)

	// ใช้ square stage แบบ concurrent 2 ชุด (สามารถต่อขยาย fan-out ได้)
	sq1 := square(nums)
	//_ = square(nums) // ไม่แนะนำให้ใช้ input เดียวกัน 2 รอบแบบนี้ (อธิบายด้านล่าง)

	// พิมพ์ผลลัพธ์จาก pipeline (เลือกช่องทางเดียวเพื่อหลีกเลี่ยง race)
	printResults(sq1)
	//printResults(sq2)

	//in := generate(1, 2, 3, 4, 5)
	//
	//// fan-out
	//c1 := square(in)
	//c2 := square(in)
	//
	//// fan-in
	//for n := range merge(c1, c2) {
	//	fmt.Println(n)
	//}
}
