package main

import (
	"fmt"
	"time"
)

func worker(channel chan int, doneChannel <- chan bool){
	for{
		select{
			case num, ok := <- channel:
				if !ok{
					return
				}
				fmt.Println(num)
			case <- doneChannel:
				return
		}
	}
}

func main(){
	doneChannel := make(chan bool)
	channel := make(chan int)

	go worker(channel, doneChannel)

	for i:=0;i<4;i++{
		channel <- i
	}

	time.Sleep(time.Second * 3)
	close(doneChannel)
}