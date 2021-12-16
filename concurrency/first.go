package main

import (
	"fmt"
	"time"
)

/*
Write a program with three functions.
One will send stuff(whatever you'd like) over a channel every one second and one will receive it and print it.
The third function will tell the other two functions to stop and return(it could be the main func) after 5 seconds.
NOTE: the program can not end until the sender and receiver have returned.
*/

func SendStuff(c chan string) {
	for {
		c <- "stuff"
		c <- "cosa"
		time.Sleep(time.Second * 1)
	}
}

func PrintStuff(c chan string) {
	for {
		msg := <-c

		fmt.Println(msg)
	}

}

func main() {
	c := make(chan string)

	go SendStuff(c)
	go PrintStuff(c)

	time.Sleep(time.Second * 5)
	/*
		c1 := make(chan string)
		c2 := make(chan string)

		go func() {
			for {
				c1 <- "from 1"
				time.Sleep(time.Second * 2)
			}
		}()

		go func() {
			for {
				c2 <- "from 2"
				time.Sleep(time.Second * 3)
			}
		}()

		go func() {
			for {
				select {
				case msg1 := <-c1:
					fmt.Println("Message 1", msg1)
				case msg2 := <-c2:
					fmt.Println("Message 2", msg2)
				case <-time.After(time.Second):
					fmt.Println("timeout")
				}
			}
		}()

		var input string
		fmt.Scanln(&input)
	*/
}

func pinger(c chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		fmt.Println("sd")
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
