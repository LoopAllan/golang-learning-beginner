package main
import "fmt"

func test() (int, string){
	return 5, "test"
}

func main(){
	x,s := test()
	fmt.Println("x =", x, ", s =", s)
}