/**
 * @Author: zhu
 * @Date: 20-4-29 下午6:00
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{3, 7, 1, 3, 6, 9, 4, 1, 8, 5, 2, 0}
	a := sort.IntSlice(i)
	fmt.Println(sort.IsSorted(a)) // false
	sort.Sort(a)
	fmt.Println("aaa",a) // [0 1 1 2 3 3 4 5 6 7 8 9]
	fmt.Println(sort.IsSorted(a), "/n") // true

	b := sort.IntSlice{3}
	fmt.Println("bbb",sort.IsSorted(b), "/n") // true

	// 更改排序行为
	c := sort.Reverse(a)
	fmt.Println(sort.IsSorted(c)) // false
	fmt.Println(c) // &{[0 1 1 2 3 3 4 5 6 7 8 9]}
	sort.Sort(c)
	fmt.Println("ccc",c) // &{[9 8 7 6 5 4 3 3 2 1 1 0]}
	fmt.Println(sort.IsSorted(c), "/n") // true

	// 再次更改排序行为
	d := sort.Reverse(c)
	fmt.Println(sort.IsSorted(d)) // false
	sort.Sort(d)
	fmt.Println(d) // &{0xc42000a3b0}
	fmt.Println(sort.IsSorted(d)) // true
	fmt.Println(d) // &{0xc42000a3b0}
}
