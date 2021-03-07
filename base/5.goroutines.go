package base

import (
	"fmt"
)

func Prints() {
	go func() {
		fmt.Println("协程执行")
	}()

	fmt.Println("执行主进程")
}
