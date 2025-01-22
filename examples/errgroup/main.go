package main

import ( 
	"golang.org/x/sync/errgroup"
	"fmt"
)

func main() {
	
	// Create a new errgroup
	var g errgroup.Group

	// Add a new goroutine to the errgroup
	g.Go(func() error {
		// Do some work
		return nil
	})

	// Add a new goroutine to the errgroup
	g.Go(func() error {
		// Do some work
		return fmt.Errorf("oops, failed")
	})

	// Wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		fmt.Println("At least one goroutine failed:", err)
	}
}

