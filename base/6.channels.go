package base

import (
	"fmt"
)

func SendMemssage(mes string) {
	ok := make(chan string)
	go func() {
		ok <- mes
	}()

	message := <-ok

	fmt.Println(message)
}
