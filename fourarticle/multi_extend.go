package main

import "fmt"

type Camera struct {

}

func (c *Camera) TakeAPictue() string {
	return "click"
}


type Phone struct {

}

func (p *Phone) Call() string {
	return "Ring Ring"
}

type CameraPhone struct {
	Camera
	Phone
}

func main() {
	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone exhibits multiple behaviors...")
	fmt.Println("It exhibits behavior of a Camera: ", cp.TakeAPictue())
	fmt.Println("It works like a Phone too: ", cp.Call())
}

