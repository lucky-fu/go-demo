package base

import "fmt"

type Student struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func CreateStudent() {
	// 结构体的初始化
	student1 := Student{"Tom", 12}

	student2 := Student{
		Name: "Jack",
		Age:  13,
	}

	student3 := &Student{
		Name: "King",
		Age:  12,
	}

	fmt.Println(student1, student2, *student3)
}
