/**
 * @Author: zhu
 * @Date: 20-5-21 下午6:49
 */
package main

import (
	"errors"
	"fmt"
	"time"
)

type Post struct {
	body string
	publishDate int64
	nest *Post
}
type Feed struct {
	length int
	start *Post
	end *Post
}

func (f *Feed)Append(newPost *Post)  {
	newPost.publishDate=time.Now().Unix()
	if f.length==0 {
		f.start=newPost
		f.end=newPost
	}else {
		lastPost:=f.end
		lastPost.nest=newPost
		f.end=newPost
	}
	f.length++
}
func main1()  {
	f:=&Feed{}
	p1:=Post{
		body:"zhu xiao is a good man",
	}
	f.Append(&p1)
	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("First: %v\n", f.start)
	p2:=Post{
		body:"but he has some shortcoming",
	}
	f.Append(&p2)

	fmt.Printf("Length: %v\n", f.length)
	fmt.Printf("First: %v\n", f.start)
	fmt.Printf("Second: %v\n", f.start.nest)
}
func (f *Feed)Remove(publishDate int64)  {
	if f.length==0 {
		fmt.Println(errors.New("Feed is empoy"))
	}
	var prviousPost *Post
	currentPost:=f.start
	for currentPost.publishDate!=publishDate {
		if currentPost.nest == nil {
			panic(errors.New("No such Post found."))
		}
		prviousPost =currentPost
		currentPost = currentPost.nest
	}
	prviousPost.nest=currentPost.nest
	f.length--
}
func (f *Feed)Insert(newPost *Post)  {
	if f.length==0 {
		f.start=newPost
	}else {
		var previousPost *Post
		currentPost :=f.start
		for currentPost.publishDate<newPost.publishDate {
			previousPost=currentPost
			currentPost=previousPost.nest
		}
		previousPost.nest=newPost
		newPost.nest=currentPost
	}
	f.length++
}
//单链表反转
//节点
type ListNode struct {
	Val int
	Next *ListNode
}
//链表
type ListAll struct {
	len int64
	pre *ListNode
	end *ListNode
}

func (l *ListAll)Append(e *ListNode)  {
	if l.len==0 {
		l.pre=e
		l.end=e
	}else {
		lastNode:=l.end
		lastNode.Next=e
		l.end=e
	}
	l.len++
}

func main()  {
	t:=&ListNode{
		Val:1,
	}
	t1:=&ListNode{
		Val:2,
	}
	t2:=&ListNode{
		Val:3,
	}
	t3:=&ListNode{
		Val:4,
	}
	t4:=&ListNode{
		Val:5,
	}
	g:=&ListAll{}
	g.Append(t)
	g.Append(t1)
	g.Append(t2)
	g.Append(t3)
	g.Append(t4)
	pr:=reverseList(t)
    fmt.Println(pr)
	for pr.Next!=nil {
		fmt.Printf("%d",pr.Next)
		pr=pr.Next
	}
}
func main2() {
	head := new(ListNode)
	head.Val = 1
	ln2 := new(ListNode)
	ln2.Val = 2
	ln3 := new(ListNode)
	ln3.Val = 3
	ln4 := new(ListNode)
	ln4.Val = 4
	ln5 := new(ListNode)
	ln5.Val = 5
	head.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln5

	pre := reverseList(head)
	fmt.Println(pre)
}

func reverseList(head *ListNode)*ListNode  {
	if head==nil || head.Next==nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur!=nil {
		fmt.Println("-----666",cur,cur.Next,pre)
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}