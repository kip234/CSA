//第一行输入N、M和T，迷宫规格N行M列，T为障碍总数
//第二行输入起点坐标Sx，Sy
//第三行输入重点坐标Ex，Ey
//第四行至第T+3行为障碍点坐标x，y
//以坐标形式输出任意一条从起点到终点的路径
package Task3

import "fmt"

//const(
//	ROAD=0//路
//	WALL=-1//墙
//	OR=-2//起点
//	DE=-3//终点
//)
//给值取的别名，方便记忆和修改
const(
	ROAD=0//路
	WALL=1//墙
	OR=2//origin-起点
	DE=3//destination-终点
	FOOT=-1//脚印
)

/*
这是一颗指向上一级的单向的树，好处在于方便开枝散叶
最后从终点节点回溯到根节点
*/
type TreeBFS struct {//记录走过的路的树的节点
	Pre *TreeBFS//上一个节点
	Coordinates [2]int//坐标
}

//object 地图对象，储存地图相关信息
type object struct {
	MAP []int//地图
	N int//行
	M int//列
	T int//障碍总数
	de [2]int//目的地坐标
	or [2]int//起点坐标
	wall [][2]int//障碍坐标
	way [][2]int//路径坐标
}

// @title function
// @description 输出寻路结果
// @auth 刘帅
// @param 无
// @return 无
func (o object)OutWay()  {
	for i:=len(o.way)-1;i>=0;i--{
		fmt.Println(o.way[i])
	}
}

// @title function
// @description 输出地图,用于编写过程中的检查,不用于功能实现
// @auth 刘帅
// @param 无
// @return 无
func (o object)OutMAP()  {
	for i:=0;i<o.N;i++ {//行
		for j:=0;j<o.M;j++{//列
			fmt.Printf("%d ",o.MAP[i*o.M+j])//输出坐标(i,j)的值
		}
		fmt.Println()//换行
	}
}

// @title function
// @description 初始化地图
// @auth 刘帅
// @param 无
// @return 无
func (o *object)Init()  {
	fmt.Scanf("%d %d %d",&o.N,&o.M,&o.T)
	//fmt.Println("=")
	fmt.Scanf("%d %d",&o.or[0],&o.or[1])//获取出发点坐标
	//fmt.Println("=")
	fmt.Scanf("%d %d",&o.de[0],&o.de[1])//获取目的地坐标
	//fmt.Println("=")
	var tmp [2]int//临时储存障碍坐标
	for i:=0;i<o.T;i++{//接收障碍坐标
		fmt.Scanf("%d %d",&tmp[0],&tmp[1])
		o.wall=append(o.wall,tmp)
	}
	//fmt.Println("接收障碍坐标完毕！")
	//填充地图
	for i:=0;i<o.M*o.N;i++{
		o.MAP=append(o.MAP,ROAD)
	}
	//fmt.Println("填充地图完毕！")
	//添加出发点与目的地
	o.MAP[o.de[0]+o.de[1]*o.M]=DE
	o.MAP[o.or[0]+o.or[1]*o.M]=OR
	//fmt.Println("添加出发点与目的地完毕！")
	//添加障碍
	for _,m:=range o.wall{
		o.MAP[m[0]+m[1]*o.M]=WALL
	}
	//fmt.Println("添加障碍完毕！")
}

//============================================================BFS系列====================================================

// @title function
// @description 产生下一层节点的列表,按上一节点又分为多个子列表
// @auth 刘帅
// @param m [][2]int 上一层的坐标列表-父节点坐标列表
// @return  ways [][][2]int 下一层的坐标列表：其中按父节点坐标又分为多个子列表
func (o *object)ProduceBFS(m [][2]int) (ways [][][2]int) { //产生下一层节点的列表,按上一节点又分为多个子列表
	for _,value:=range m {
		var Tmp [][2]int
		x, y :=value[0],value[1]
		if x > 0 && (o.MAP[y*o.M+x-1] == ROAD || o.MAP[y*o.M+x-1] == DE) { //向左走
			tmp := [2]int{x - 1, y}
			Tmp = append(Tmp, tmp)
			o.MAP[y*o.M+x-1]=FOOT
		}
		if x < o.M-1 && (o.MAP[y*o.M+x+1] == ROAD || o.MAP[y*o.M+x+1] == DE) { //向右走
			tmp := [2]int{x + 1, y}
			Tmp = append(Tmp, tmp)
			o.MAP[y*o.M+x+1]=FOOT
		}
		if y > 0 && (o.MAP[(y-1)*o.M+x] == ROAD || o.MAP[(y-1)*o.M+x] == DE) { //向上走
			tmp := [2]int{x, y - 1}
			Tmp = append(Tmp, tmp)
			o.MAP[(y-1)*o.M+x]=FOOT
		}
		if y < o.N-1 && (o.MAP[(y+1)*o.M+x] == ROAD || o.MAP[(y+1)*o.M+x] == DE) { //向下走
			tmp := [2]int{x, y + 1}
			Tmp = append(Tmp, tmp)
			o.MAP[(y+1)*o.M+x]=FOOT
		}
		ways=append(ways,Tmp)
	}
	return
}

// @title function
// @description 建立用于储存路线的树
// @auth 刘帅
// @param be *TreeBFS 根节点
// @param dest [2]int 目的地坐标
// @return  re bool 成功创建(有终点)
// @return de *TreeBFS 指向终点所在节点的指针
func (o *object)BuildTreeBFS(be *TreeBFS,dest [2]int) (re bool,de *TreeBFS) {
	way := o.ProduceBFS([][2]int{be.Coordinates})
	var rootList = []*TreeBFS{be}
	for {

		if len(way) == 0 {
			re = false
			de = nil
			return
		}
		//fmt.Println(len(way))
		var treeList [][2]int
		var rootListTmp []*TreeBFS
		for i, value1 := range way {
			for _, value2 := range value1 {
				temp := new(TreeBFS)
				treeList=append(treeList,value2)//最后的结果是：所有的子列表合并成一个列表
				rootListTmp=append(rootListTmp,temp)//记录声明的节点，作为下一层的父节点
				temp.Pre = rootList[i]
				temp.Coordinates = value2
				if value2 == dest {
					re = true
					de = temp
					return
				}
			}

		}
		//fmt.Println(treeList)
		//fmt.Println(way)
		way=o.ProduceBFS(treeList)
		rootList=rootListTmp
	}
}

// @title function
// @description 负责寻路
// @auth 刘帅
// @param x int 起点x坐标
// @param y int 起点y坐标
// @return  bool 成功寻路
func (o *object)ExploreBFS(x,y int) bool {
	tmp := [2]int{x, y}       //以数组的形式储存当前坐标，方便记录
	tree:=new(TreeBFS)
	tree.Coordinates=tmp
	isok,re:=o.BuildTreeBFS(tree,[2]int{o.de[0],o.de[1]})
	if isok {
		for re!=nil {
			o.way=append(o.way,re.Coordinates)
			re=re.Pre
		}
		return true
	}
	return false
}

//===========================================================DFS系列=====================================================

// @title function
// @description 负责寻路
// @auth 刘帅
// @param x int 当前x坐标
// @param y int 当前y坐标
// @return  bool 成功寻路
func (o *object)ExploreDFS(x,y int) bool {
	tmp:=[2]int{x,y}//以数组的形式储存当前坐标，方便记录
	// y*o.M+x 为该点在切片中的索引值
	if o.MAP[y*o.M+x]==DE {//到达终点
		o.way=append(o.way,tmp)//记录位置-已到达终点
		return true
	}
	if x>0 && (o.MAP[y*o.M+x-1]==ROAD||o.MAP[y*o.M+x-1]==DE) {//向左走
		o.MAP[y*o.M+x]=FOOT //做标记-留下脚印
		if o.ExploreDFS(x-1,y){//前往下一处
			o.way=append(o.way,tmp)//记录位置-成功到达终点
			return true
		}
	}
	if x<o.M-1 && (o.MAP[y*o.M+x+1]==ROAD||o.MAP[y*o.M+x+1]==DE) {//向右走
		o.MAP[y*o.M+x]=FOOT //做标记-留下脚印
		if o.ExploreDFS(x+1,y){//前往下一处
			o.way=append(o.way,tmp)//记录位置-成功到达终点
			return true
		}
	}
	if y>0 && (o.MAP[(y-1)*o.M+x]==ROAD||o.MAP[(y-1)*o.M+x]==DE) {//向上走
		o.MAP[y*o.M+x]=FOOT //做标记-留下脚印
		if o.ExploreDFS(x,y-1){//前往下一处
			o.way=append(o.way,tmp)//记录位置-成功到达终点
			return true
		}
	}
	if y<o.N-1 && (o.MAP[(y+1)*o.M+x]==ROAD||o.MAP[(y+1)*o.M+x]==DE) {//向下走
		o.MAP[y*o.M+x]=FOOT //做标记-留下脚印
		if o.ExploreDFS(x,y+1){//前往下一处
			o.way=append(o.way,tmp)//记录位置-成功到达终点
			return true
		}
	}
	return false//都走投无路了还没到终点
}



// @title function
// @description 封装功能的函数
// @auth 刘帅
// @param 无
// @return 无
func Task()  {
	var m object
	m.Init()
	//m.OutMAP()
	m.ExploreBFS(m.or[0],m.or[1])
	//m.ExploreDFS(m.or[0],m.or[1])
	//m.OutMAP()
	m.OutWay()
	//fmt.Println(m)
}

/*
  |0 1 2 3 4 5 6
--|-------------> X
0 |0 0 0 0 0 0 0
1 |0 1 0 1 0 1 0
2 |0 1 0 1 0 1 0
3 |0 1 0 1 1 1 0
4 |1 0 0 0 0 1 1
5 |1 0 1 1 0 0 0
6 |0 0 0 0 0 1 0
  |
  Y
 */

//input:
/*
7 7 17
0 0 6 6
0 4
0 5
1 1
1 2
1 3
2 5
3 1
3 2
3 3
3 5
4 3
5 1
5 2
5 3
5 4
5 6
6 4
*/

//output
/*

*/