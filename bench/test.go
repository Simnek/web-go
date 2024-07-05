package bench

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func FastFor() {
	var wg sync.WaitGroup

	for i := 1; i <= 1000000; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			// Simulate some work
			x := 0
			for j := 0; j < 1000; j++ {
				x += j
			}
			fmt.Println("Goroutine", index, "value of x:", x)
		}(i)
	}

	wg.Wait()
}

func SlowFor() {
	for i := 1; i <= 1000000; i++ {
		// Simulate some work
		x := 0
		for j := 0; j < 1000; j++ {
			x += j
		}
		fmt.Println("Goroutine", i, "value of x:", x)
	}
}

func BenchmarkFastFor(b *testing.B) {
	iterations := b.N

	startTime := time.Now()
	for i := 0; i < iterations; i++ {
		FastFor()
	}
	endTime := time.Now()
	fmt.Println("FastFor took", endTime.Sub(startTime), "for", iterations, "iterations")

	startTime = time.Now()
	for i := 0; i < iterations; i++ {
		SlowFor()
	}
	endTime = time.Now()
	fmt.Println("SlowFor took", endTime.Sub(startTime), "for", iterations, "iterations")
}
