package main

import "fmt"

func main() {


	s := "good by"
	var p *string  = &s
	*p = "ciao"
	fmt.Printf("p%p\n", p)
	fmt.Printf("*p%s\n", *p)
	fmt.Printf("s:%s\n", s)


}


