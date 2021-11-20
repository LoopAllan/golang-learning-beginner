package main
import "fmt"

func main(){
	// 判斷式練習
	var b bool
	fmt.Print("輸入布林值")
	fmt.Scanln(&b)
	if b{
		fmt.Println("Go")
	} else {
		fmt.Println("Not Go")
	}

	// 簡易情境練習： ATM 檢測使用者輸入的金額是否適當
	var money int
	fmt.Print("請問要領多少錢：")
	fmt.Scanln(&money)
	if money < 100{
		fmt.Println("Too Few!")
	} else if money <= 1e6{
		fmt.Println("OK!")
	} else {
		fmt.Println("Insufficient Funds!")
	}
	fmt.Println("執行完畢！")
}