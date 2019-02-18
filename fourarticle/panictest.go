package main

import "fmt"

func main() {
	fmt.Println("starting the progrm")
	panic("啊server error occurred: stoping the program")
	fmt.Println("Ending the program")

}
