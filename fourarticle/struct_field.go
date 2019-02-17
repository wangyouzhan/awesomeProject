package main

type struct1 struct {
	i1 int
	f1  float32
	str string
}

type File struct {
	fd int
	name string
}

func NewFile(fd int, name string) *File{
	if fd < 0 {
		return nil
	}
	return &File{fd,name}
}

//func main() {
//
//
//	f := NewFile(12, "./test.txt")
//
//
//	//ms := new(struct1)
//	//ms.i1 = 10
//	//ms.f1 = 15.0
//	//ms.str = "Chris"
//	//
//	//fmt.Printf("The int is : %d\n", ms.i1)
//	//fmt.Printf("The float is : %f\n", ms.f1)
//	//fmt.Printf("The string is : %s\n", ms.str)
//
//
//}
