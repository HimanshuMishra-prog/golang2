package main

import (
	"fmt"
	"sync"
)

func calculate(sum chan int, i int,
	wg *sync.WaitGroup) {
	var temp int
	for ; i < 1000; i = i + 5 {
		temp += i * i
		fmt.Println(temp)
	}
	sum <- temp
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	sum := make(chan int)
	sum1 := make(chan int)
	sum2 := make(chan int)
	sum3 := make(chan int)
	sum4 := make(chan int)

	wg.Add(5)
	go calculate(sum, 1, &wg)
	go calculate(sum1, 2, &wg)
	go calculate(sum2, 3, &wg)
	go calculate(sum3, 4, &wg)
	go calculate(sum4, 5, &wg)

	/*for i := 1; i <= 5; i++ {
		fmt.Println("correct3")

		fmt.Println("correct4")
		wg.Add(1)
		go calculate(sum, i, &wg, &m)

		fmt.Println("correct5")
	}*/
	fmt.Println("correct")
	total := <-sum + <-sum1 + <-sum2 + <-sum3 + <-sum4
	wg.Wait()
	fmt.Println("The sum is  :", total)

}

//commented out
/*
func calculate(sum chan int, i int,
	wg *sync.WaitGroup, m *sync.Mutex) {

	var temp int
	for ; i < 1000; i = i + 5 {
		temp += i * i


	}
	m.Lock()
	sum <- temp
	m.Unlock()
	wg.Done()

}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	sum := make(chan int)


	/*for i := 1; i <= 5; i++ {
		wg.Add(1)
		go calculate(sum, i, &wg, &m)
	}

	fmt.Println("The sum is  :", <-sum)
	wg.Wait()

	//fmt.Println("The sum is  :", <-sum)

}
*/
