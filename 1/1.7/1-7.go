package main

import (
	"fmt"
	"unsafe"
)

func main() {
	func0()
	func1()
	func2()
}

func func0() {
	type data struct {
		a int
	}
	var d = data{1234}
	var p *data
	p = &d
	fmt.Printf("%p, %v\n", p, p.a) // 直接用指针访问目标对象成员，无须转换。
}

func func1() {
	x := 0x12345678
	p := unsafe.Pointer(&x) // *int -> Pointer
	n := (*[4]byte)(p)      // Pointer -> *[4]byte
	for i := 0; i < len(n); i++ {
		fmt.Printf("%X ", n[i])
	}
}

func func2() {
	d := struct {
		s string
		x int
	}{"abc", 100}
	p := uintptr(unsafe.Pointer(&d)) // *struct -> Pointer -> uintptr
	p += unsafe.Offsetof(d.x)        // uintptr + offset
	p2 := unsafe.Pointer(p)          // uintptr -> Pointer
	px := (*int)(p2)                 // Pointer -> *int
	*px = 200                        // d.x = 200
	fmt.Printf("%#v\n", d)
}
