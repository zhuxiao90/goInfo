/**
 * @Author: zhu
 * @Date: 20-4-23 上午10:28
 */
package workjob

type Worker struct {
	JobQueue chan Job
}

func NewWorker()Worker  {
	return Worker{JobQueue:make(chan Job)}
}
func (w Worker)Run(wq chan chan Job)  {
	go func() {
		for  {
			wq <- w.JobQueue
			select {
			case job:=<-w.JobQueue:
				job.Do()
			}
		}
	}()
}