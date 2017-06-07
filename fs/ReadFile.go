package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	var st syscall.Stat_t
	if err := syscall.Stat("/tmp/sysstat.go", &st); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("mtime: %d\n", st.Mtimespec.Sec)
	fmt.Printf("ctime: %d\n", st.Ctimespec.Sec)
}