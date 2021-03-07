package main

import (
	"fmt"

	"github.com/demo/base"
)

func main() {
	base.CreateStudent()

	react := base.MewReact(1, 3)
	println(react.GetArea())

	base.SendMemssage("哈哈哈")

	fmt.Println("哈哈")
}
