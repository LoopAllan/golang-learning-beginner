package main
import "fmt"

type employee struct{
	name string
	age int
	gender string
	job_position string
}

func create_employee(name string, age int, gender string, job_position string) employee{
	e := employee{name, age, gender, job_position}
	return e
}

func print_employee(e employee){
	fmt.Println(e.name, e.age, e.gender, e.job_position)
}

func main(){
	var e1 employee = create_employee("小明", 37, "男性", "程式設計師")
	print_employee(e1)
	e2 := create_employee("高登", 52, "男性", "CEO")
	print_employee(e2)
	e2.age = 55
	print_employee(e2)
}