package main

import (
	"os"
	"syscall"
)

const ENABLE_VIRTUAL_TERMINAL_PROCESSING = 0x0004

func init() {
	stdout := syscall.Handle(os.Stdout.Fd())

	var originalMode uint32
	syscall.GetConsoleMode(stdout, &originalMode)
	originalMode |= ENABLE_VIRTUAL_TERMINAL_PROCESSING

	syscall.MustLoadDLL("kernel32").MustFindProc("SetConsoleMode").Call(uintptr(stdout), uintptr(originalMode))
}
