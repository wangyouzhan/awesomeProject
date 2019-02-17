package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int
	innerS

}


func main() {

	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 4
	outer.in1 = 9
	outer.in2 = 8
	fmt.Printf("%d\n", outer.in2)


	//使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{6, 10}}
	fmt.Println("outer2 is : ", outer2)

}
