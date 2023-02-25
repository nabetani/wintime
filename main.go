package main

import (
	"fmt"
	"time"
)

func show(msg string, ticks []time.Time) {
	borders := []float64{ //
		0.37, 0.42, 0.49, 0.56, 0.65, 0.75, 0.87, 1.00, 1.15, 1.33, 1.54, 1.78,
		2.05, 2.37, 2.74, 3.16, 3.65, 4.22, 4.87, 5.62, 6.49, 7.50, 8.66, 10.00,
		11.55, 13.34, 15.40, 17.78, 20.54, 23.71, //
		1e100}
	slots := make([]int, len(borders)+1)
	for i := 0; i < len(ticks)-1; i++ {
		nano := ticks[i+1].Sub(ticks[i]).Nanoseconds()
		ms := float64(nano) * 1e-6
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
	ticks := []time.Time{}
	ticker := time.NewTicker(time.Millisecond)
	for range ticker.C {
		ticks = append(ticks, time.Now())
		if 5000 < len(ticks) {
			break
		}
	}
	show(msg, ticks)
}

func main() {
	test("before")
	for _, i := range []uint32{1, 15, 16, 256} {
		timeBeginPeriod(i)
		test(fmt.Sprint("tbp:", i))
		timeEndPeriod(i)
	}
}
