package main

import (
	"fmt"
	"time"
)

var globalCh chan bool

func t1(ch1, ch2 chan int) {
	//get the value
	for {
		a := <-ch1
		fmt.Printf("this is 1 : a : %v\n", a)
		if a > 5 {
			globalCh <- true
			return
		}
		//send the value
		ch2 <- a + 2
		//time.Sleep(time.Second * 1)
	}
}

func t2(ch1, ch2 chan int) {
	for {
		a := <-ch2
		fmt.Printf("this is 2 : a : %v\n", a)
		if a > 50 {
			globalCh <- true
			return
		}

		ch1 <- a + 1
		//time.Sleep(time.Second * 1)
	}
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go t1(ch1, ch2)
	go t2(ch1, ch2)
	ch1 <- 0

	select {
	case <-globalCh:
		fmt.Println("done1")

	case <-time.After(time.Second * 10):
		fmt.Println("done2")
	}
	fmt.Println("done done")
	time.Sleep(time.Second * 500)
}
