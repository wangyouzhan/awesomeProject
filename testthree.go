package main

import (
	"fmt"
	"strconv"
)

type TZ int



func main() {

	//var str string = "This is an example of a string"
	//fmt.Printf("T/F? %s has prefix %s", str, "Th")
	//fmt.Printf("\n%t\n", strings.HasPrefix(str, "Th"))

	//var str string = "Hello, how is it going, Hugo?"
	//var manyG = "gggggggg"


	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("%d\n", strconv.IntSize)



	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is : %d\n", an)

	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)






	//fmt.Printf("Number of H's in %s is: ", str)
	//fmt.Printf("%d\n", strings.Count(str, "H"))
	//fmt.Printf("%s", strings.Repeat(str, 3))
	//fmt.Printf("\n%s", strings.ToUpper(str))

	//s1 := strings.Fields(str)
	//fmt.Printf("%v", s1)
	//for _, va1 := range s1{
	//	fmt.Printf("%s - ", va1)
	//}




}
