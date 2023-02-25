//go:build windows

package main

import (
	"log"
	"syscall"
	"unsafe"
)

var kernel32 = func() *syscall.DLL {
	p, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		log.Panicln(err)
	}
	return p
}()

func kernel32Proc(name string) *syscall.Proc {
	proc, err := kernel32.FindProc(name)
	if err != nil {
		log.Panicln(err)
	}
	return proc
}

var pQueryPerformanceCounter = kernel32Proc("QueryPerformanceCounter")
var pQueryPerformanceFrequency = kernel32Proc("QueryPerformanceFrequency")

func queryPerformanceCounter() uint64 {
	v := uint64(0)
	pQueryPerformanceCounter.Call(uintptr(unsafe.Pointer(&v)))
	return v
}
func queryPerformanceFrequency() uint64 {
	v := uint64(0)
	pQueryPerformanceFrequency.Call(uintptr(unsafe.Pointer(&v)))
	return v
}
