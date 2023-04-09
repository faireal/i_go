package main

import (
	"fmt"
	"math"
)

func main() {
	// 十进制
	var a = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010

	// 八进制
	var b = 077
	fmt.Printf("%o \n", b)

	// 十六进制
	var c = 0xff
	fmt.Printf("%x \n", c)
	fmt.Printf("%X \n", c)

	fmt.Println(c)

	// 浮点数
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	// 字符串
	var s1 = "faireal"
	fmt.Println(s1)

	s2 := `line1 
line2
line3
`
	fmt.Println(s2)

	// byte rune
	s3 := "faireal 饭饭"
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%v-%c \n", s3[i], s3[i])
	}

	for _, str := range s3 {
		fmt.Printf("%v-%c \n", str, str)
	}
	s4 := []rune(s3)
	s5 := []byte(s3)
	fmt.Println(s4)
	fmt.Println(s5)
}
