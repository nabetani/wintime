package main

import (
	"fmt"
	"time"
)

func show(msg string, ticks []uint64) {
	tickMS := 1000.0 / float64(queryPerformanceFrequency())
	borders := []float64{ //
		0.12, 0.13, 0.15, 0.18, 0.21, 0.24, 0.27, 0.32,
		0.37, 0.42, 0.49, 0.56, 0.65, 0.75, 0.87, 1.00, 1.15, 1.33, 1.54, 1.78,
		2.05, 2.37, 2.74, 3.16, 3.65, 4.22, 4.87, 5.62, 6.49, 7.50, 8.66, 10.00,
		11.55, 13.34, 15.40, 17.78, 20.54, 23.71, //
		1e100}
	slots := make([]int, len(borders)+1)
	for i := 0; i < len(ticks)-1; i++ {
		ms := float64(ticks[i+1]-ticks[i]) * tickMS
		for pos := 0; ; pos++ {
			if ms < borders[pos] {
				slots[pos]++
				break
			}
		}
	}
	fmt.Print(msg)
	for _, v := range slots {
		fmt.Printf(",%d", v)
	}
	fmt.Println()
}

func test(msg string) {
	const size = 5000
	ticks := make([]uint64, 0, size)
	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()
	for range ticker.C {
		ticks = append(ticks, queryPerformanceCounter())
		if size < len(ticks) {
			break
		}
	}
	show(msg, ticks)
}

func main() {
	test("before")
	for _, i := range []uint32{1, 15, 16, 256, 1024} {
		timeBeginPeriod(i)
		test(fmt.Sprint("tbp:", i))
		timeEndPeriod(i)
	}
}
