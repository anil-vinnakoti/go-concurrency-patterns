package main

import "fmt"


// Each value walks through the pipeline step-by-step, blocking at each stage until the next stage is ready.
func pipeline() {
	generatedNumsChannel := generator(2,5,8)
	squaredNumsChannel := squareNums(generatedNumsChannel)
	for squredNum := range squaredNumsChannel{
		fmt.Println("squared num:", squredNum)
	}
	
}

func generator(nums... int)<-chan int{
	numsChannel := make(chan int)
	go func(){
		for _, num := range nums{
			numsChannel <- num
		}
		close(numsChannel)
	}()

	return  numsChannel
}

func squareNums(ch <-chan int) <-chan int{
	squaredChannel := make(chan int)
	go func ()  {
		for num := range ch{
			squaredChannel <- num * num
		}
		close(squaredChannel)
	}()
	return squaredChannel
}