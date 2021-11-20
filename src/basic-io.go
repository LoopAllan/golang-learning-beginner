package main
import "fmt"
func main(){
	// 基本輸出練習 fmt.Println(資料, 資料, 資料, ...)
	fmt.Println(3, "Hello", true)

	// 基本輸入練習 fmt.Scanln(&{變數名稱}, &{變數名稱})
	// &變數名稱: 取得變數的指標 (Pointer)
	var x int
	fmt.Print("輸入整數：")
	fmt.Scanln(&x)	// 為給值預設為０
	fmt.Println("x =", x)

	// 整合練習：輸入兩個整數，並輸出相乘的結果
	var x1 int
	var y1 int
	fmt.Print("輸入第一個數字：")
	fmt.Scanln(&x1)
	fmt.Print("輸入第二個數字：")
	fmt.Scanln(&y1)
	var result int = x1 + y1
	fmt.Println("x + y =", result)

	// 同時輸入兩個變數
	var x2 int
	var y2 int
	fmt.Print("輸入兩個數字(用空格隔開)：")
	fmt.Scanln(&x2, &y2)
	var result2 int = x2 + y2
	fmt.Println("x + y =", result2)
}