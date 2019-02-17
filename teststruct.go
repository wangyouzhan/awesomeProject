package main

import (
	"awesomeProject/five"
	"awesomeProject/structPack"
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool "An important answer"
	field2 string  "The name of the thing"
	field3 int "How much there are"
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

func main() {

	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}



	wrong := five.NewMatrix(12)
	fmt.Printf("%v\n", wrong)

	struct1 := new(structPack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 10.9

	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
	fmt.Printf("Mf1 = %f\n", struct1.Mf1)

}
