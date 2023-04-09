package _4_goroutine

import (
	"fmt"
	"sync"
)

var wg = new(sync.WaitGroup)

func f1() {
	fmt.Println("hello world")
}

func f2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello world")

}

func f3(wg *sync.WaitGroup, n int) {
	defer wg.Done()
	fmt.Printf("hello world : %d \n", n)

}
