/**
 * @Author: zhu
 * @Date: 20-5-25 下午5:16
 */
package main

import "fmt"

//定义二叉树节点
type Node struct {
	value int
	Left,Right *Node
} 
//打印二叉树
func (node *Node)Print()  {
	fmt.Println(node.value)
}
//设置value值
func (node *Node)SetValue(v int)  {
	if node==nil {
		fmt.Println("不要给空指针赋值啊")
		return
	}
	node.value=v
}
//初始化节点
func InitNode(v int)*Node{
	return &Node{value:v}
}
func main()  {
	root:=Node{value:3}
	root.Left=&Node{}
	root.Left.SetValue(0)
	root.Left.Right=InitNode(2)
	root.Right=&Node{5,nil,nil}
	root.Right.Left=InitNode(4)
	//root.BreathRange()
	tf:=root.Layers()
	fmt.Println("---rty",tf)
}
//前序遍历
func (node *Node)FrontRange()  {
	if node==nil {
		return
	}
	node.Print()
	node.Left.FrontRange()
	node.Right.FrontRange()
}
//中序遍历
func (node *Node)MidRange()  {
	if node==nil {
		return
	}
	node.Left.MidRange()
	node.Print()
	node.Right.MidRange()
}
//中序遍历
func (node *Node)AftRange()  {
	if node==nil {
		return
	}
	node.Left.AftRange()
	node.Right.AftRange()
	node.Print()

}
//层次遍历
func (node *Node)BreathRange()  {
	if node==nil {
		return
	}
	var result []int
	nodes :=[]*Node{node}
	for len(nodes)>0 {
		curNode:=nodes[0]
		nodes =nodes[1:]
		result=append(result,curNode.value)
		if curNode.Left!=nil {
			nodes=append(nodes,curNode.Left)
		}
		if curNode.Right!=nil {
			nodes=append(nodes,curNode.Right)
		}
	}
	for _, value := range result {
		fmt.Print(value," ")
	}
}
//层数实现（递归）
func (node *Node)Layers()int  {
	if node==nil {
		return 0
	}
	leftLayers := node.Left.Layers()
	rightLayers := node.Right.Layers()
	if leftLayers > rightLayers {
		return leftLayers + 1
	} else {
		return rightLayers + 1
	}
}
//层数(非递归实现)
//借助队列，在进行按层遍历时，记录遍历的层数即可
func (node *Node) LayersByQueue() int {
	if node == nil {
		return 0
	}
	layers:=0
	nodes:=[]*Node{node}
	for len(nodes)>0 {
		layers++
		size:=len(nodes)
		count :=0
		for count<size {
			count++
			curNode:=nodes[0]
			nodes=nodes[1:]
			if curNode.Left != nil {
				nodes = append(nodes, curNode.Left)
			}
			if curNode.Right != nil {
				nodes = append(nodes, curNode.Right)
			}
		}
	}
	return layers
}
