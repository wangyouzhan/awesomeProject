package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	outputFile, outputError := os.OpenFile("./fourarticle/products.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError  != nil{
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}

	defer outputFile.Close()


	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"


	for i:= 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()

}
