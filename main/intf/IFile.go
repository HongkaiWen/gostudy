package intf

type IFile interface {
	Read() (n int, err error)
	Write() (n int, err error)
}
