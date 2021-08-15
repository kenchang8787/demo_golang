/*
Stringer

fmt 套件中定義的 Stringer 是最普遍的接口之一

type Stringer interface {
    String() string
}

Stringer 是一個可以用字符串描述自己的類型，fmt 套件 (或其他套件) 都通過此接口來打印值。
*/
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 定義實作 Stringer 接口的方法，可以影響 fmt.Println 方法打印的值
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
