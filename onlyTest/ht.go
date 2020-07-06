/**
 * @Author: zhu
 * @Date: 20-4-23 下午10:01
 */
package main

import (
	"fmt"
	"sort"
)

func main()  {
	infos = []int{14, 1, 42, 18, 9, 23, 24, 5, 10, 29,99}
	//SuPai(0, len(infos )-1)
	ErFun(infos,29)
	//ChaPai(infos)
	fmt.Println(infos)
}
var infos []int
func SuPai(left,right int)  {
	i:=left
	j:=right
	ba:=infos[i]
	for i<j  {
		for infos[j]>ba {
			j--
		}
		for infos[i]<ba {
			i++
		}
		infos[i],infos[j]=infos[j],infos[i]
	}
	if left>=right {
		return
	}
	SuPai(left, i-1)
	if i!=len(infos)-1 {
		SuPai(i+1,right)
	}

}
func ChaPai(ss []int)  {
	n:=len(ss)
	if n<2 {
		return
	}
	for i:=1;i<n ;i++  {
		for j:=i;j>0&&ss[j]<ss[j-1] ;j--  {
			ss[j],ss[j-1]=ss[j-1],ss[j]
		}

	}
	fmt.Println(ss)
}
func XiEr(s []int)  {
	n:=len(s)
	h:=1
	for h<n/3 {
		h=3*h+1
	}
	for h>=1 {
		//将数组变为间隔h的有序元素
		for i:=h;i<n ;i++  {
			for j:=i;j>=h&&s[j]<s[j-h] ;j-=h  {
				s[j],s[j-h]=s[j-h],s[j]
			}
		}
		h/=3
	}
}
func XuaZe(m []int,star int)  {
	if star==len(m) {
		return
	}
	minIdx := star
	minVal  :=m[star]
	for i:=star+1;i<len(m) ;i++  {
		if m[i]<minVal {
			minIdx,minVal=i,m[i]
		}
	}
	m[star],m[minIdx]=m[minIdx],m[star]
	XuaZe(m,star+1)
}
func MaoPao(n []int)  {
	for i:=0;i<len(n) ;i++  {
		for j:=i+1;j<len(n) ;j++  {
			if n[i]>n[j] {
				n[i],n[j]=n[j],n[i]
			}
		}
	}
}
func ErFun(e []int,ox int)  {
	sort.Ints(e)
	/*fmt.Println("---jk",e)
	hl:= sort.Reverse(sort.IntSlice(e))
	sort.Sort(hl)
	fmt.Println("---rrr",hl)*/
    s,f:=0,len(e)-1
	for s<=f {
		mid:=(s+f)/2
		if e[mid] == ox {
			fmt.Println("---aaa",mid)
			break
		}
		if e[mid] < ox {
			s=mid+1
		}
		if e[mid] > ox {
			f=mid-1
		}
	}
}