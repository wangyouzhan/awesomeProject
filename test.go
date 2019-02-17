package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)


var (
	home = os.Getenv("HOME")
	user = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

func init() {
	if user == ""{
		log.Fatal("$USER not set")
	}
	if home == ""{
		home = "/home/" + user
	}
	if gopath == ""{
		gopath = home + "/go"
	}
	fmt.Printf(gopath)
	fmt.Println()
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}

func main()  {
	//fmt.Println(stringutil.Reverse("!oG, olleHoooo"))
	//fmt.Println(stringutil.ReverseTwo("wang youzhan!"))

	//for i := 0; i < 5; i++ {
	//	defer fmt.Printf("%d ", i)
	//}

	type T struct {
		a int
		b float64
		c string }

	var timeZone = map[string]int{
		"UTC": 0*60*60,
		"EST": -5*60*60,
		"CST": -6*60*60,
		"MST": -7*60*60,
		"PST": -8*60*60,
	}

	t := &T{7,  -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", timeZone)




}