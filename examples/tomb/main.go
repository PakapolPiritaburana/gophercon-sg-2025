package main

import (
	"gopkg.in/tomb.v2"
	"time"
	"fmt"
)

func main() {
	var t tomb.Tomb

	t.Go(func() error {
		for {
			select {
			case <-t.Dying():
				fmt.Println("shutting down because of", t.Err())
				return t.Err()
			default:
				// do something
			}
		}
	})

	t.Go(func() error {
		time.AfterFunc(2*time.Second, func() {
			t.Kill(fmt.Errorf("timeout"))
		})
		return nil
	})

	fmt.Println("waiting for goroutines to finish")
	t.Wait()
	fmt.Println("all goroutines finished")
}
