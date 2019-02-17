package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("setaability of v: ", v.CanSet())
	v = reflect.ValueOf(&x)
	fmt.Println("type of v: ", v.Type())
	fmt.Println("setability of v: ", v.CanSet())
	v = v.Elem()
	fmt.Println("The Elem of v is : ", v)
	fmt.Println("setableility of v:", v.CanSet())

	v.SetFloat(3.1415926)
	fmt.Println(v.Interface())
	fmt.Println(v)

}
