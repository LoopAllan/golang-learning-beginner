package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func pow(x, n, lim float64) float64 {
	// if 條件句結束前可以做變數宣告，作用範圍於 if-else 之內
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func os() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func date() {
	fmt.Println("When's Saturday")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}

func dura() {
	t := time.Now()
	fmt.Println("現在時間是 ", t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	// switch 應用，作業系統判斷
	os()
	// switch 應用，多久以後禮拜六
	date()
	// switch 應用，判斷時間點打招呼
	dura()
	// 計算整數多次平方並限制範圍
	fmt.Println("３的平方為不超過10為", pow(3, 2, 10))
	fmt.Println("３的３次方不超過20為", pow(3, 3, 20))
	// 判斷式練習
	var b bool
	fmt.Print("輸入布林值")
	fmt.Scanln(&b)
	if b {
		fmt.Println("Go")
	} else {
		fmt.Println("Not Go")
	}

	// 簡易情境練習： ATM 檢測使用者輸入的金額是否適當
	var money int
	fmt.Print("請問要領多少錢：")
	fmt.Scanln(&money)
	if money < 100 {
		fmt.Println("Too Few!")
	} else if money <= 1e6 {
		fmt.Println("OK!")
	} else {
		fmt.Println("Insufficient Funds!")
	}
	fmt.Println("執行完畢！")
}
