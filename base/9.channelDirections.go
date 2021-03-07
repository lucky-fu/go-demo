package base

func ping(pings chan<- string, mes string) {
	pings <- mes
}

func pong(pings <-chan string, pongs chan<- string) {
	mes := <-pings
	pongs <- mes
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "hah")
	pong(pings, pongs)
}
