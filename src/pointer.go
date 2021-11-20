package main
import "fmt"

func main(){
	var x int = 3
	fmt.Println("x =", x)
	fmt.Println("&x =", &x)
	var xPtr *int = &x
	fmt.Println("xPtr =", xPtr)
	fmt.Println("&xPtr =", &xPtr)
	fmt.Println("*xPtr =", *xPtr)

	var word string = "Hello"
	fmt.Println("word =", word)
	var wordPtr *string = &word
	fmt.Println("wordPtr =", wordPtr)
	fmt.Println("*wordPtr =", *wordPtr)
	word2 := *wordPtr
	fmt.Println("word2 =", word2)
}