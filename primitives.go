package main

import "fmt"

func main(){
	channelOne := make(chan string)
	channelTwo := make(chan string)

	go func(){
		channelTwo <- "Two"
		// close(channelTwo)
	}()

	go func(){
		channelOne <- "One"
		// close(channelOne)
	}()

	select{
		case msgFromOne := <-channelOne:
			fmt.Println(msgFromOne)
		case msgFromTwo := <-channelTwo:
			fmt.Println(msgFromTwo)
	}
		
}