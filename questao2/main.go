package main

import (
	"fmt"
	"sync"
	"time"
  "math/rand"
)

var qtdVezes []int
var n int
 
func worker(waiting1 *sync.WaitGroup, waiting2 *sync.WaitGroup, id int) {
  	var t = rand.Intn(6)
  	var vezSorteada int

  	fmt.Printf("Primeira fase -> goroutine com o id(%v) dormindo por %v segundos\n", id, t)
	time.Sleep(time.Duration(t) * time.Second)

  	var s = rand.Intn(11)	  
  	qtdVezes[id-1] = s

	fmt.Printf("Primeira fase -> goroutine com o id(%v) e valor aleatorio = %v\n", id, s)

  	waiting1.Done()
	waiting1.Wait()

  	fmt.Println("Segunda fase -> Iniciando goroutine com o id:", id)

  	if id == 1 {
		vezSorteada = qtdVezes[n-1] 
		time.Sleep(time.Duration(vezSorteada) * time.Second)
	} else {
		vezSorteada = qtdVezes[id-2] 
		time.Sleep(time.Duration(vezSorteada) * time.Second)
	}

    fmt.Printf("Segunda fase -> goroutine com o id(%v) dormindo por %v segundos\n", id, vezSorteada)
  	fmt.Printf("Segunda fase -> goroutine com o id(%v) finalizado após %v segundos\n", id, vezSorteada)

  	waiting2.Done()
  
}

func main() {
	fmt.Println("Digite o valor de n desejado:")
	fmt.Scan(&n)

	var waiting1 sync.WaitGroup
	var waiting2 sync.WaitGroup

	qtdVezes = make([]int, n)

	for i := 1; i <= n; i++ {
		fmt.Println("Primeira fase -> Iniciando goroutine com o id:", i)
		waiting1.Add(1)
    waiting2.Add(1)
		go worker(&waiting1, &waiting2, i)
	}

	waiting2.Wait()
  	fmt.Printf("Segunda fase -> Concluido depois de %v goroutines\n", n)
  	fmt.Println("Finalizado")
}