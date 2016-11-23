package main

import (
	"fmt"
	"mytest"
	"win_tool"
)

func main() {
	fmt.Printf("Hello,world. Sqrt(2) = %v\n", mytest.Mytest(2))

	win_tool.Wait_input()
}
