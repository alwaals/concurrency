package main

import (
	//"bytes"
	"fmt"
	"strings"

	//"strings"
	"time"
)

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
	}()

	return ch
}

//  func main() {
// 	timeStart := time.Now()
// 	go func(){
// 		<-worker()
// 		<-worker()
// 	}()

//		println(int(time.Since(timeStart).Seconds())) // 3 or 6 ?
//	 }
func main() {

	v := 5
	p := &v
	println(*p, &p, p)

	changePointer(p, strings.Builder{})
	fmt.Println(*p)

	var abs string = "swetha1"
	for _, v := range abs {
		if v >= '0' && v <= '9' {
			fmt.Println("It has capital letters")
		}
	}
	var b1 strings.Builder
	b1.WriteString("ABC")
	b2 := b1
	fmt.Println(b2.Len())    // 3
	fmt.Println(b2.String()) // ABC
	b2.Reset()
	b2.WriteString("DEF")
	fmt.Println(b2.String()) // DEF
}

func changePointer(p *int, a strings.Builder) {
	println(p)
	v := 3
	p = &v
	println(*p)
	var b1 strings.Builder
	b1 = a
	b1.WriteString("added")
	fmt.Println(a.String(), b1.String())
}

// func main() {
// 	ch := make(chan string)
// 	ch2 := make(chan struct{})
// 	var wg sync.WaitGroup

// 	wg.Add(3)
// 	go func() {
// 		ch <- "Hello this is 1st goroutine"
// 		defer wg.Done()
// 	}()

// 	go func() {
// 		close(ch2)
// 		defer wg.Done()
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		for {
// 			select {
// 			case msg := <-ch:
// 				fmt.Println(msg)
// 			case <-ch2:
// 				return
// 			}
// 		}
// 	}()
// 	wg.Wait()
// }

// Atomic usage

// var sharedRes int32

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func() {
// 		for i := 1; i <= 3; i++ {
// 			atomic.AddInt32(&sharedRes, 1)
// 		}
// 		wg.Done()
// 	}()
// 	go func() {
// 		for i := 1; i <= 3; i++ {
// 			atomic.AddInt32(&sharedRes, 1)
// 		}
// 		wg.Done()
// 	}()
// 	wg.Wait()
// 	fmt.Println(sharedRes)
// }
