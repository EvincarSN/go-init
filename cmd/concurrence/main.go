package main

// <- antes da variável - "esvasiar"
// <- depois da variável - "encher"

import (
	"fmt"
	"time"
)

func processando() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

}

// T1
func main() {
	// go processando() //T2
	// go processando() //T3
	// processando()    //T1

	canal := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			canal <- i //T2
			fmt.Println("Jogou o canal", i)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			canal <- i //T2
			fmt.Println("Jogou o canal", i)
		}
	}()

	// for x := range canal {
	// 	fmt.Println("Recebeu do canal", x)
	// 	time.Sleep(time.Second)
	// }

	go worker(canal, 1)
	go worker(canal, 2)
	go worker(canal, 3)
	worker(canal, 4)
}

func worker(canal chan int, workerID int) {
	for {
		fmt.Println("Recebeu do canal", <-canal, "no worker", workerID)
		time.Sleep(time.Second)
	}
}
