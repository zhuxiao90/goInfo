/**
 * @Author: zhu
 * @Date: 20-5-22 下午2:16
 */
package main

import (
	"container/list"
	"fmt"
)

//节点结构
type DNode struct {
	data interface{}
	prev *DNode
	next *DNode
}
//链表数据结构
type DList struct {
	size uint64
	head *DNode
	tail *DNode
}
//链表初始化
func InitList()(list *DList)  {
	list=new(DList)
	list.size=0
	list.head=nil
	list.tail=nil
	return
}
func (dList *DList)Move(e,at *DNode)  {

}
func (dList *DList)GetSize()uint64  {
     return dList.size
}

func main() {
	//初始化双链表
	link := list.New()
    //循环赋值
	for i := 0; i <= 10; i++ {
		link.PushBack(i)
	}
    //遍历查询
	for p := link.Front(); p != link.Back(); p = p.Next() {
		fmt.Println("Number", p.Value)
	}

}

