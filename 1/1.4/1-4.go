package main

import (
	"fmt"
)

const x, y int = 1, 2     // 多常量初始化
const s = "Hello, World!" // 类型推断
const (
	Sunday    = iota // 0
	Monday           // 1，通常省略后续⾏行表达式。
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

func main() {
	const x = "xxx" // 未使⽤用局部常量不会引发编译错误。
	fmt.Println(Saturday)

}
