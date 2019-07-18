package Awsfuzzer

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/AndreZiviani/aws-fuzzer/AwsFuzzerVFS"
)

const (
	mfdCloexec  = 0x0001
	memfdCreate = 319
)

func runFromMemory(displayName string, filePath string) {
	fdName := "" // *string cannot be initialized
	fd, _, errno := syscall.Syscall(memfdCreate, uintptr(unsafe.Pointer(&fdName)), uintptr(mfdCloexec), 0)

	if errno != 0 {
		panic(errno)
	}

	buffer, err := AwsFuzzerVFS.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	_, err = syscall.Write(int(fd), buffer)
	if err != nil {
		panic(err)
	}

	fdPath := fmt.Sprintf("/proc/self/fd/%d", fd)
	err = syscall.Exec(fdPath, []string{displayName}, nil)

	if err != nil {
		panic(err)
	}
}

func main() {
	runFromMemory("fzf", "bin/fzf")
}
