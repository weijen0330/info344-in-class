package main

import (
	"fmt"
	"math/rand"
	"time"
)

//someLongFunc is a function that might
//take a while to complete, so we want
//to run it on its own go routine
// channel of ints
func someLongFunc(ch chan int) {
	r := rand.Intn(2000)
	d := time.Duration(r)
	time.Sleep(time.Millisecond * d)

	// writing r into the channel. Will block if channel is full
	ch <- r
}

func main() {
	//TODO:
	//create a channel and call
	//someLongFunc() on a go routine
	//passing the channel so that
	//someLongFunc() can communicate
	//its results
	rand.Seed(time.Now().UnixNano())
	fmt.Println("starting long-running func...")
	// unbuffered, accept 1 int at a time
	ch := make(chan int)
	go someLongFunc(ch)

	// read out the result
	result := <-ch
	fmt.Printf("result was %d\n", result)
}
