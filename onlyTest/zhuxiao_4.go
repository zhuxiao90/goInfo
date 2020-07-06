/**
 * @Author: zhu
 * @Date: 20-5-20 下午6:06
 */
package main

import (
	"fmt"
)

//channel是一个数据类型，引用类型，主要用来解决go程的同步问题以及go程之间数据共享（数据传递）的问题
//goroutine是通过通信来共享内存，而不是共享内存来通信
//引用类型的零值就都是nil，比如slice，map，interface，channel，指针
//channel关闭后，再往channel里面写数据会panic，但是可以继续读数据，对于nil channel，无论收发都会被阻塞
//通过close（chan）来关闭channel，使用if data, ok := <-c; ok，来判断channel是否有某值，data := range c遍历channel
//可以将 channel 隐式转换为单向队列，只收或只发，不能将单向 channel 转换为普通 channel。c:=make(chan int, 3)  var send chan<- int = c // send-only
//突然想起了一句话，linux系统中一切皆文件
//缓冲器的作用是：解耦，处理并发，缓存
//select case case语句里必须是一个IO操作，监听i/o操作，用于监听channel上的数据流动
//死锁：两个或两个以上进程在执行过程中，由于竞争资源或者彼此通信而造成的一种阻塞的现象。
//sync.RWMutex 读写锁 读共享写独占（RLock共享Lock独占）
type OrderInfo struct {
	id int
}
func main(){
	c:=make(chan OrderInfo,10)
	d:=make(chan int)
	go cou(c)
	go printer(c,d)
	<-d
	fmt.Println("done")
}
func cou(out chan OrderInfo)  {
	defer close(out)
	for i:=0;i<10 ;i++  {
		order:=OrderInfo{id:i}
		out<-order
	}
}
func printer(in chan OrderInfo,h chan int)  {
	for nm:=range in{
		fmt.Println("order id is :",nm.id)
	}
	h<-9
}
func fb(c,q chan int)  {
	x,y:=1,1
	for {
		select {
		case c<-x:
			x,y=y,x+y
		case <-q:
			fmt.Println("qqq")
			return
		}
	}
}
/*func main()  {
	c:=make(chan int)
	q:=make(chan int)
	go func() {
		for {
			select {
			case v:=<-c:
				fmt.Println(v)
			case <-time.After(5*time.Second):
				fmt.Println("qqq")
				q<-999
				return
			}
		}
	}()
	c<-666
	<-q
}*/
/*func main() {
	ch := make(chan int)
	  		// I'm blocked because there is no channel read yet.
	fmt.Println("send")
	go func() {
		<-ch  		// I will never be called for the main routine is blocked!
		fmt.Println("received")
	}()
	ch <- 1
	fmt.Println("over")
}*/