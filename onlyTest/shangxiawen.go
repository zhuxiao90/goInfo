/**
 * @Author: zhu
 * @Date: 2020/6/15 下午6:33
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)
//context包可以提供一个请求从API请求边界到各goroutine的请求域数据传递、取消信号及截至时间等能力。
//Context是安全的，可被多个goroutine同时使用。一个Context可以传给多个goroutine，而且可以给所有这些goroutine发取消信号。
//context包提供从已有Context衍生新的Context的能力。这样即可形成一个Context树，当父Context取消时，所有从其衍生出来的子Context亦会被取消。
//Background是所有Context树的根，其永远不会被取消。
//Deadline — 返回 context.Context 被取消的时间，也就是完成工作的截止日期；
//Done — 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消之后关闭，多次调用 Done 方法会返回同一个 Channel；
//Err — 返回 context.Context 结束的原因，它只会在 Done 返回的 Channel 被关闭时才会返回非空的值；
//如果 context.Context 被取消，会返回 Canceled 错误；
//如果 context.Context 超时，会返回 DeadlineExceeded 错误；
//Value — 从 context.Context 中获取键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法可以用来传递请求特定的数据；
func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// monitor
		go func() {
			for range time.Tick(time.Second) {
				select {
				case <-r.Context().Done():
					fmt.Println("req is outgoing")
					return
				default:
					fmt.Println("req is processing")
				}
			}
		}()

		// assume req processing takes 3s
		time.Sleep(3 * time.Second)
		w.Write([]byte("hello"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
