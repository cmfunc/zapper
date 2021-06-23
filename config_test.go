package zaper

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestConfig(t *testing.T)  {
	a:=Config{}

	// Alignof 一个类型对齐值
	// Sizeof 类型占用字节数
	fmt.Println(unsafe.Alignof(a))
	fmt.Println(unsafe.Sizeof(a))

	fmt.Println(unsafe.Alignof(a.ConsoleWriter))
	fmt.Println(unsafe.Sizeof(a.ConsoleWriter))

	fmt.Println(unsafe.Alignof(a.ConsoleEncoder))
	fmt.Println(unsafe.Sizeof(a.ConsoleEncoder))

	fmt.Println(
		unsafe.Alignof(a.LogFile)+
		unsafe.Alignof(a.HandlingLevel)+
		unsafe.Alignof(a.ConsoleWriter)*3+
		unsafe.Alignof(a.ConsoleWriter)*3,
	)
}