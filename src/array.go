package main

import "fmt"

func main() {
	var numbers [3]int
	numbers[0] = 15
	numbers[1] = 30
	numbers[2] = 60
	// numbers[3] = 100 // 不能超過陣列長度
	fmt.Println(numbers)
	fmt.Println(numbers[1])
	fmt.Println("numbers 長度：", len(numbers))
	// 範例，使用者入３位學生資料，並印出平均及總和
	fmt.Println("輪詢陣列中資料")
	var grades [3]int
	var sum int = 0
	for i := 0; i < len(grades); i++ {
		fmt.Print("請輸入第", i, "位學生成績：")
		fmt.Scanln(&grades[i])
		sum += grades[i]
	}
	fmt.Println("成績平均：", sum/len(grades))
	fmt.Println("成績總和：", sum)
}
