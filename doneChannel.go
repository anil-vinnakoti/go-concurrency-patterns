package main

import (
	"fmt"
	"time"
)

func someWorker(channel chan int, doneChannel <- chan bool){
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

func doneChannel(){
	doneChannel := make(chan bool)
	channel := make(chan int)

	go someWorker(channel, doneChannel)

	for i:=0;i<4;i++{
		channel <- i
	}

	time.Sleep(time.Second * 1)
	close(doneChannel)
}