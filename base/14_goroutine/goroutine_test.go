package _4_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestG1(t *testing.T) {
	f1()
	fmt.Println("main")
}

func TestG2(t *testing.T) {
	go f1()
	fmt.Println("main") // 输出: main   f1未执行原因 主程序退出导致未执行
}

func TestG3(t *testing.T) {
	go f1()
	fmt.Println("main")
	time.Sleep(time.Millisecond * 30) // 输出: main  hello world   由于主程序阻塞了一段时间 f1能够执行下去 所有
}

func TestG4(t *testing.T) {
	wg.Add(1)
	go f2(wg)
	fmt.Println("main")
	wg.Wait() // 等待所有并发都执行
}

func TestG5(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f3(wg, i)
	}
	fmt.Println("main")
	wg.Wait()
}

func TestG6(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)

	}
}

// 无缓冲channel 下面这段代码会报错 fatal error: all goroutines are asleep - deadlock!
// 如果对一个无缓冲通道执行接收操作时，没有任何向通道中发送值的操作那么也会导致接收操作阻塞
func TestC1(t *testing.T) {
	ch := make(chan struct{})
	<-ch
	fmt.Println("done")
}

// 无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段
func TestC11(t *testing.T) {
	ch := make(chan struct{})
	ch <- struct{}{}
	fmt.Println("done")
}

// 发送值给channel
func TestC2(t *testing.T) {
	ch := make(chan struct{})
	go func(ch chan struct{}) {
		time.Sleep(time.Millisecond * 30)
		ch <- struct{}{}
	}(ch)
	<-ch

	fmt.Println("done")
}

func TestC22(t *testing.T) {
	ch := make(chan struct{})
	go func(ch chan struct{}) {
		time.Sleep(time.Millisecond * 30)
		<-ch
	}(ch)
	ch <- struct{}{}

	fmt.Println("done")
}

// 有缓冲channel
func TestC3(t *testing.T) {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	fmt.Println("done")
}

// 死锁, 一直会等待接收通道
func TestC33(t *testing.T) {
	ch := make(chan struct{}, 1)
	<-ch
	fmt.Println("done")
}
