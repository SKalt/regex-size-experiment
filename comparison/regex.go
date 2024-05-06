package main

import (
	"fmt"
	"regexp"
	"runtime"
	"unsafe"
)

func main() {
	var pre, post runtime.MemStats
	runtime.ReadMemStats(&pre)
	re := regexp.MustCompile(`0.0.`) // allocations seem to be 760 + 16*len... ish.
	// The complexity of the regex seems to have an effect on the number of allocations.
	runtime.ReadMemStats(&post)
	fmt.Printf("sizeof: %d\n", unsafe.Sizeof(re))
	fmt.Printf("alloc: %d\n", post.Alloc-pre.Alloc)
	fmt.Printf("heap: %d\n", post.HeapAlloc-pre.HeapAlloc)
}
