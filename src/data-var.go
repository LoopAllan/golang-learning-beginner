package main
import "fmt"
func main(){
	// 輸出訊息到終端： fmt.Println(data1, data2, ...)
	fmt.Println(3) // 整數 int
	fmt.Println(3.1415)	// 浮點數 float64
	fmt.Println("測試中文")	// 字串 string
	fmt.Println(true) // 布林值 bool
	fmt.Println('a') // 字符 rune
	// 變數的使用
	var x int // 宣告變數 x
	x = 4 // 把資料指派給 x
	fmt.Println("x =", x)
	x = 120
	fmt.Println("x =", x)
	/* 	
		資料型態指派錯誤範例：
			x = "Hello"	// 將出錯，資料型態不相符:
			# command-line-arguments
			./data-var.go:16:4: cannot use "Hello" (type untyped string) as type int in assignment
	*/
	var f float64 = 3.1415
	fmt.Println("f =", f)
	var s string = "哈囉！Golang!"
	fmt.Println("s =", s)
	var b bool = true
	fmt.Println("b =", b)
	var c rune = 'b'
	fmt.Println("c =", c)
}