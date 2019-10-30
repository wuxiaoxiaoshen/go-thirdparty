package _interface

type Action interface {
	Do(interface{})
	String() string
	Run(interface{})
	Close()
}
