package main

import "fmt"

func main() {
	var a []string
	var b = []int{}
	var c = []bool{true}

	fmt.Println(a == nil) // a 只是声明了 未初始化所以为 true
	fmt.Println(b == nil)
	fmt.Println(c == nil)

	a = append(a, "a")
	fmt.Println(a)

	// 切片的长度和容量
	// 切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。

	m := [5]int{1, 2, 3, 4, 5}
	s := m[1:3] // s[2,3]
	fmt.Printf("s:%v len(s):%v cap(s):%v \n", s, len(s), cap(s))

	// a[2:] ==  a[2:len(a)]
	// a[:3] == a[0:3]
	// a[:] == a[0:len(a)]
	// 0 <= low <= high <= len(a)
	s2 := s[3:4] // s2= [5]
	fmt.Println(s2)

	m1 := make([]int, 2, 10)
	fmt.Println(m1)

	// 切片的赋值拷贝
	t1 := make([]int, 2, 10)
	t2 := t1 // 共用底层数组
	t11 := make([]int, len(t1))
	t1[0] = 10
	copy(t11, t1) // 深度拷贝
	t1[1] = 100
	fmt.Println(t1, t2, t11)

	// 两个切片合并
	var t3 = []int{1, 2, 3}
	var t4 = []int{4, 5, 6}
	t3 = append(t3, t4...)
	fmt.Println(t3)

}
