package main
import "fmt"

func add(xPtr *int){
	fmt.Println("add()")
	*xPtr += 10
	fmt.Println(*xPtr)
}

func main(){
	var x int = 10
	xPtr := &x
	fmt.Println("x =", x)
	add(xPtr) // Pass By Pointer
	fmt.Println("x =", x)
	fmt.Println("*xPtr =", *xPtr)

	// 範例：和使用者要求輸入
	var msg string
	var msgPtr *string = &msg
	fmt.Print("輸入：")
	// fmt.Scanln(&msg) 改為下方方式
	fmt.Scanln(msgPtr)
	fmt.Println("輸出：", msg)
}