package main

import (
	"fmt"
	"sync"
	"time"
  "math/rand"
)

func worker(waiting1 *sync.WaitGroup, id int) {

	defer waiting1.Done()
  	var t = rand.Intn(6)

	fmt.Printf("goroutine id(%v): Started waiting %v seconds\n", id, t)

	time.Sleep(time.Duration(t) * time.Second)

	fmt.Printf("goroutine id(%v): Finished after %v seconds\n", id, t)
}

func main() {
	var n int

  	fmt.Println("Digite o Valor de n desejado:")
	fmt.Scan(&n)

	var waiting1 sync.WaitGroup

	for i := 1; i <= n; i++ {
		fmt.Println("Starting goroutine com id:", i)
		waiting1.Add(1)
		go worker(&waiting1, i)
	}

	waiting1.Wait()


	fmt.Printf("Goroutine Completo com n = %v\n", n)
}