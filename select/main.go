package main

func main() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	go func() {
		anotherChannel <- "more data"
	}()

	select {
	case msg1 := <-myChannel:
		println("Received:", msg1)
	case msg2 := <-anotherChannel:
		println("Received:", msg2)
	}
}
