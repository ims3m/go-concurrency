package main

// To communicate info between goroutines
func main() {

	myChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	msg := <-myChannel
	println(msg)

}
