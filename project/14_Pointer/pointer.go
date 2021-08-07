package main

import "fmt"

func main() {

	fmt.Println("\nthis is call by value")

	var x = 1
	var y = foo1(x)
	fmt.Printf("x: %d\ny: %d", x, y)

	q := [3]int{1, 2, 3}
	r := foo2(q)
	fmt.Printf("\nq: %v\nr: %d", q, r)

	fmt.Println("\n\nthis is call by reference")

	var a = 1
	var b = fooo1(&a)
	fmt.Printf("a: %d\nb: %d", a, b)
	// slice
	m := []int{1, 2, 3}
	n := fooo2(m)
	fmt.Printf("\nm: %v\nn: %d", m, n)
}

func foo1(y int) int {
	y = y + 1
	return y
}

func foo2(y [3]int) int {
	var z = y[0] + 1
	return z
}

func fooo1(a *int) int {
	*a = *a + 1
	return *a
}

// use a slice rather than a array (a slice is a window of array with array's pointer)
func fooo2(sli []int) int {
	sli[0] = sli[0] + 1
	return sli[0]
}
