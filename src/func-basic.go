package main
import "fmt"

func add(a float64, b float64) float64{
	return a+b;
}

func sub(a float64, b float64) float64{
	return a-b;
}

func division(a float64, b float64) float64{
	return a/b;
}

func mutiply(a float64, b float64) float64{
	return a*b;
}

func say(msg string){
	fmt.Println(msg)
}

func sum(from float64, to float64, step float64) float64{
	var result, x float64
	for x=from; x<=to; x+=step {
		result += x
	}
	return result
}

func main(){
	var a, b, c float64
	fmt.Println("Step1. 加法，請輸入兩個數字：")
	fmt.Scanln(&a, &b)
	fmt.Println("結果：", add(a, b))

	fmt.Println("Step2. 減法，請輸入兩個數字：")
	fmt.Scanln(&a, &b)
	fmt.Println("結果：", sub(a, b))

	fmt.Println("Step3. 除法，請輸入兩個數字：")
	fmt.Scanln(&a, &b)
	fmt.Println("結果：", division(a, b))

	fmt.Println("Step4. 乘法，請輸入兩個數字：")
	fmt.Scanln(&a, &b)
	fmt.Println("結果：", mutiply(a, b))

	fmt.Println("Step5. 加總，請輸入兩個數字：")
	fmt.Scanln(&a, &b, &c)
	fmt.Println("結果：", sum(a, b, c))

	var msg string
	fmt.Println("Step6. 說話，請輸入一句話：")
	fmt.Scanln(&msg)
	say(msg)
}