/**
 * @Author: zhu
 * @Date: 20-4-23 上午10:56
 */
package workjob

import "fmt"

type WorkerPool struct {
	Workerlen   int
	JobQueue    chan Job
	WorkerQueue chan chan Job
}

func NewWorkerPool(workerlen int) *WorkerPool {
	return &WorkerPool{
		Workerlen:   workerlen,
		JobQueue:    make(chan Job),
		WorkerQueue: make(chan chan Job, workerlen),
	}
}

func (wp *WorkerPool)Run()  {
	fmt.Println("初始化worker")
	for i:=0;i<wp.Workerlen ;i++  {
		worker:=NewWorker()
		worker.Run(wp.WorkerQueue)
	}
	//循环获取可用的worker，往worker中写job
	go func() {
		for  {
			select {
			case job:=<-wp.JobQueue:
				worker:=<-wp.WorkerQueue
				worker<-job
			}
		}
	}()
}
