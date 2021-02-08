//手动输入一串英文单词，以"nil"为结尾，随后输入第一个单词的首字母，然后程序依次不重复输出直至结尾为止。
package Task2

import (
	"fmt"
	"strings"
)

////用于储存排列结果的树
//type Tree struct {
//	Pre *Tree
//	Data int
//	Next []*Tree
//}
//
//func Out(t Tree)  {
//	for len(t.Next)!=0{
//		for _,n:=range t.Next{
//			fmt.Print(n.Data)
//		}
//	}
//	fmt.Println()
//}

//data 数据对象，储存单词及状态
type data struct {
	Contain []string//内容->单词
	Usable []bool//单词对应的状态，是否可用
}

// @title function
// @description 负责单词的输入
// @auth 刘帅
// @param stop string 暂停输入的标志
// @return 无
func (d *data)Input(stop string)  {
	var tmp string//临时储存单词
	fmt.Scanf("%s",&tmp)//第一次输入
	for !strings.EqualFold(tmp,stop){//是否为停止标志
		d.Contain=append(d.Contain,tmp)//记录单词
		d.Usable=append(d.Usable,true)//添加状态
		fmt.Scanf("%s",&tmp)//刷新
	}
}

// @title function
// @description 输出接龙结果
// @auth 刘帅
// @param head byte 第一个单词的首字母
// @return 无
func (d *data)Out(head byte)  {
	flag:=true//是否存在符合要求的
	for flag{
		flag=false
		for i,m:=range d.Contain{//遍历寻找以head开头并且没用过的单词
			if m[0]==head && d.Usable[i]{//找到了
				fmt.Printf("%s ",m)//输出
				d.Usable[i]=false//标记已经用过
				head=m[len(m)-1]//刷新
				flag=true
				break
			}
		}

	}
}
// @title function
// @description 封装功能的函数
// @auth 刘帅
// @param 无
// @return 无
func Task()  {
	var a data
	a.Input("nil")
	//var tree Task2.Tree
	//Task2.Sort('b',&tree,a)
	var tmp byte
	fmt.Scanf("%c",&tmp)
	a.Out(tmp)
	//fmt.Println(tree)
}

////对单词进行排列
//func Sort(m byte,tree *Tree,d Data) {
//	tmpD:=d
//	for i,c:=range tmpD.Contain{//寻找以m开头并没有使用的单词
//		if c[0]==m && tmpD.Usable[i]!=false{//找到了
//			tmp:=new(Tree)
//			tmp.Data=i//记录节点
//			tree.Next=append(tree.Next,tmp)
//			tmpD.Usable[i]=false//标记
//			Sort(tmpD.Contain[i][len(tmpD.Contain[i])-1],tmp,d)
//		}
//	}
//}