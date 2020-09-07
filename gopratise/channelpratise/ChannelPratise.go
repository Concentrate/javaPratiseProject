package main

import (
	"fmt"
	"sync"
	"time"
)

func enterAndLeaveCall(tag string) string {
	fmt.Printf("\nenter  %v<<<<<\n", tag)
	return fmt.Sprintf("\nleave  %v>>>>>>\n", tag)
}

func selectAndDefaultUse() {
	tmpChan := make(chan int, 3)
	go func() {
		for i := 0; i < 10; i++ {
			tmpChan <- i
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			select {
			case a1 := <-tmpChan:
				fmt.Printf("select case value is %v \n", a1)

			case <-time.After(300 * time.Millisecond):
				fmt.Println("after duration ")

				//default:
				//	fmt.Printf("default case\n ")
				//	time.Sleep(300 * time.Millisecond)

			}
		}
	}()

}

func rangeUseChannelConsume() {
	defer fmt.Println(enterAndLeaveCall("rangeUseChannelConsume"))

	waitDone := sync.WaitGroup{}
	tmpChan := make(chan int, 3)
	waitDone.Add(10)
	go func() {
		for i := 0; i < 10; i++ {
			tmpChan <- i
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for a1 := range tmpChan {
			fmt.Println("range use ", a1)
			waitDone.Done()
		}
	}()

	//for a1 := range tmpChan {
	//	fmt.Println("range use ", a1)
	//	waitDone.Done()
	//}
	waitDone.Wait()
}

func main() {
	//withoutBuffer()
	//withBuffer()
	//selectAndDefaultUse()
	rangeUseChannelConsume()
	time.Sleep(time.Minute)
}

func withBuffer() {
	defer fmt.Print(enterAndLeaveCall("withBuffer "))
	produceAndConsumer(3)

}
func withoutBuffer() {
	defer fmt.Print(enterAndLeaveCall("withoutBuffer "))
	produceAndConsumer(0)
}

func produceAndConsumer(buffer int) {
	tmpChan := make(chan int, buffer)

	go func() {
		for i := 0; i < 3; i++ {
			tmpChan <- i
			fmt.Printf("produce %v\n", i)
		}
	}()

	consuFun := func(tag string) {
		fmt.Printf("tag %v, and consume %v\n", tag, <-tmpChan)
	}

	go consuFun("c1")
	go consuFun("c2")
	time.Sleep(4 * time.Second)
}
