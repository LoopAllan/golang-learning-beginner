package main
import "fmt"

func main(){
	// break
	fmt.Println("break")
	var x int = 0
	for x < 5 {
		if x == 3 {
			break;
		}
		fmt.Println(x)
		x++
	}
	// continue
	fmt.Println("continue")
	x = 0
	for x=0; x<5; x++ {
		if x%2 == 0 {continue}
		fmt.Println(x)
	}
	// 範例：持續讓使用者輸入，直到使用者輸入０為止
	var sum int
	fmt.Println("請輸入輸字，將於輸入０後輸出總和：")
	for true{
		var input int = -1
		fmt.Scanln(&input)
		if input == 0 {break}
		sum += input
	}
	fmt.Println("總和為：", sum)
}