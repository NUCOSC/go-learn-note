package main

import "fmt"

func main() {
	s := "Hello, World!"
	s1 := s[:5]  // Hello
	s2 := s[7:]  // World!
	s3 := s[1:5] // ello
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	func2()
	func3()
}

func func2() {
	s := "abcdc"
	bs := []byte(s)
	bs[1] = 'C'
	println(string(bs))
	println(s)
	u := "电脑"
	us := []rune(u)
	us[1] = '话'
	println(string(us))
	println(u)
}

func func3() {
	s := "abc汉字"
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%c,", s[i])
	}
	fmt.Println()
	for _, r := range s { // rune
		fmt.Printf("%c,", r)
	}
}
