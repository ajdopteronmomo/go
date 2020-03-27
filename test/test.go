package main

import "fmt"

type Stringer interface {
	String() string
}

type Printer interface {
	String() string
	Print()
}

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("%d ,%s", self.id, self.name)
}

func (self *User) Print() {
	fmt.Println(self.String())
}

func main() {
	// var o interface{} = &User{1, "Tom"}

	//1----
	// if i, ok := o.(fmt.Stringer); ok {
	// 	fmt.Println(i)
	// }

	// u := o.(*User)
	// // u := o.(User)
	// fmt.Println(u)

	//2----
	// switch v := o.(type) {
	// case nil:
	// 	fmt.Println("nil")
	// case fmt.Stringer:
	// 	fmt.Println(v)
	// case func() string:
	// 	fmt.Println(v())
	// case *User:
	// 	fmt.Println("%d,%s\n", v.id, v.name)
	// default:
	// 	fmt.Println("unknow")
	// }

	//3----
	var o Printer = &User{1, "Tom"}
	var s Stringer = o
	fmt.Println(s.String())
}
