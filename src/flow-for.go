package main

import "fmt"

func main() {
	// 無限迴圈
	// for true {
	// 	fmt.Println("Hello!")
	// }

	// 迴圈練習： 1+2+3+4+...+50
	var x, y int
	for x = 1; x <= 50; x++ {
		fmt.Print(y, "+", x)
		y += x
		fmt.Println("=", y)
	}
	fmt.Println("y =", y)

	// while in golang
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
