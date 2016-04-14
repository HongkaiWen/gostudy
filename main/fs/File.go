package fs

import "fmt"

type File struct {
}

func (f *File) Read(buf byte[])(n int, err error) {
fmt.Println("read method")
}

func (f *File) Write(buf byte[])(n int, err error) {
fmt.Println("read method")
}
