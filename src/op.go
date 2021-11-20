package main
import "fmt"
func main(){
	// 算術運算: +, -, *, /
	fmt.Println("算術運算: +, -, *, /")
	var x int
	x = 3 * 3 + 10
	fmt.Println("x = 3 * 3 + 10 =", x)
	// 指定運算: =, += ,-=, *=, /=
	fmt.Println("指定運算: =, += ,-=, *=, /=")
	x = x + 5
	fmt.Println("x = x + 5 =", x)
	x += 5
	fmt.Println("x += 5 =", x)
	x -= 5
	fmt.Println("x -= 5 =", x)
	x *= 5
	fmt.Println("x *= 5 =", x)
	x /= 5
	fmt.Println("x /= 5 =", x)
	// 單元運算: ++, --
	fmt.Println("單元運算: ++, --")
	x++
	fmt.Println("x++ =", x)
	x--
	fmt.Println("x-- =", x)
	// 比較運算: >, <, >=, <=, ==
	fmt.Println("比較運算: >, <, >=, <=, ==")
	var result bool = 4 > 3
	fmt.Println(result)
	result = 4 < 3
	fmt.Println("result = 4 < 3 =", result)
	result = 4 >= 4
	fmt.Println("result = 4 >= 4 =", result)
	result = 4 <= 4
	fmt.Println("result = 4 <= 4 =", result)
	result = 4 == 3
	fmt.Println("result = 4 == 3 =", result)
	// 邏輯運算: !, &&, ||
	fmt.Println("邏輯運算: !, &&, ||")
	result = !false
	fmt.Println("result = !false =", result)
	result = !true
	fmt.Println("result = !true =", result)
	result = true && false
	fmt.Println("result = true && false =", result)
	result = true || false
	fmt.Println("result = true || false =", result)
}