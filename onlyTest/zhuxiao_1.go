/**
 * @Author: zhu
 * @Date: 20-5-14 上午10:43
 */
package main

import (
	"fmt"
	"strings"
)

//切片由指针，长度，容量三部分构成，其中指针只想以一个slice元素对应的地城数组元素地址，但不一定是第一个，比如s[2:]指向的就是底层数组下标为2的地址
//说到底，指针只是引用数组，而非数组指针或者指针数组，通过引用，实现变长方案
//s[low:high:max] len=high-low cap=max-low
//append单个元素，或者append少量的多个元素，这里的少量指double之后的容量能容纳，这样就会走以下扩容流程，不足1024，双倍扩容，超过1024的，1.25倍扩容
//若是append多个元素，且double后的容量不能容纳，直接使用预估的容量
//此外，以上两个分支得到新容量后，均需要根据slice的类型size，算出新的容量所需的内存情况capmem，然后再进行capmem向上取整，得到新的所需内存，除上类型size，得到真正的最终容量,作为新的slice的容量
func main()  {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := data[8:]  	//{8, 9}
	s2 := data[:5] 	//{0, 1, 2, 3, 4}
	copy(s2, s1)
	fmt.Println(s2)
	fmt.Println(data)
	fmt.Printf("%v\n", strings.HasPrefix("World Cup.png", "World"))
}