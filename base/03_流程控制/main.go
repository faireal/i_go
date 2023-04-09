package main

import (
	"fmt"
)

func main() {
	// is else 分支
	func(score int) {
		if score >= 90 {
			fmt.Println("A")
		} else if score >= 80 {
			fmt.Println("B")
		} else {
			fmt.Println("c")
		}

	}(65)

	// for 循环1
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
	// for 循环2
	func() {
		i := 10
		for i >= 1 {
			fmt.Println(i)
			i--
		}
	}()

	// swich
	func(score int) {
		switch score {
		case 1, 11, 111:
			fmt.Println(1)
		case 2, 22, 222:
			fmt.Println(2)
		case 3:
			fmt.Println(3)
		case 4:
			fmt.Println(4)
		default:
			fmt.Println(5)
		}
	}(6)
}
