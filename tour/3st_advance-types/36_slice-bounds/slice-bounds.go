/*
切片的默認行為

在進行切片時，可利用默認行為來定義範圍，
切片的下界默認值為 0，上界則是該切片的長度。

var a [10]int
a[0:10]
a[:10]
a[0:]
a[:]
故對於數組 a，以上的切片方式都是相同的。
*/
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
