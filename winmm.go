//go:build windows

package main

import (
	"log"
	"syscall"
)

var winmm = func() *syscall.DLL {
	p, err := syscall.LoadDLL("winmm.dll")
	if err != nil {
		log.Panicln(err)
	}
	return p
}()

func winmmProc(name string) *syscall.Proc {
	proc, err := winmm.FindProc(name)
	if err != nil {
		log.Panicln(err)
	}
	return proc
}

var pTimeBeginPeriod = winmmProc("timeBeginPeriod")
var pTimeEndPeriod = winmmProc("timeEndPeriod")

func timeBeginPeriod(period uint32) uint32 {
	pTimeBeginPeriod.Call(uintptr(period))
	return 0
}
func timeEndPeriod(period uint32) uint32 {
	pTimeEndPeriod.Call(uintptr(period))
	return 0
}
