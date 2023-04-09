package main

import "fmt"

func main() {
	var a = [3]int{1, 2, 3}
	var b = [2]string{"a", "b"}
	var c = [2]string{}
	fmt.Println(a, b, c)

	// 二位数组
	var d = [2][2]string{{"a", "b"}, {"c", "d"}}
	fmt.Println(d)
}
