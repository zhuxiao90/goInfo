/**
 * @Author: zhu
 * @Date: 20-4-14 下午3:42
 */
package main

import (
	"fmt"
	"onlyTest/workjob"
	"runtime"
	"time"
)

type Score struct {
	Num int
}

func (s *Score)Do()  {
	fmt.Println("num:",s.Num)
	time.Sleep(1*time.Second)
}

func main()  {
	num:=100
	//debug.SetMaxThreads(num+1000)//设置最大线程数
	//注册工作池，传入任务
	//参数1 woker并发个数
	p:=workjob.NewWorkerPool(num)
	p.Run()
	datanum := 100 * 100 * 100 * 100
	go func() {
		for i:=1;i<datanum ;i++  {
			sc:=&Score{Num:i}
            p.JobQueue<-sc
		}
	}()
	for  {
		fmt.Println("runtime.NumGoroutine():",runtime.NumGoroutine())
		time.Sleep(2*time.Second)
	}
}
