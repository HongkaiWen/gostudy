package intf

type Query interface {
	Query() ([]*Result, error)
}
