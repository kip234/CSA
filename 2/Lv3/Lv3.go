//这里有一个结构体，表示一个哔哩哔哩的用户
//我现在需要一个表示视频详情的结构体（请自行设计，可以参考哔哩哔哩的），并且能实现
//点赞
//收藏
//投币
//一键三连
//这几个方法。
//而且还需实现一个发布视频的函数，入参为作者名，视频名，返回一个视频结构体
package Lv3

type Author struct {
	Name string //名字
	VIP bool //是否是高贵的带会员
	Icon string //头像
	Signature string //签名
	Focus int //关注人数
}

type Video struct {
	Author
	like int
	collect int
	coins int
}

//点赞
func (v *Video)Like()  {
	v.like+=1
}

//收藏
func (v *Video)Collect()  {
	v.collect+=1
}

//投币
func (v *Video)InsertCoin()  {
	v.coins+=1
}

//一键三连
func (v *Video)OneKeyTriple()  {
	v.like+=1
	v.collect+=1
	v.coins+=1
}

func ReleaseVideos (AuthorName,VideoName string) (newVideo Video) {
	newVideo=Video{
		Author{
			AuthorName,
			false,
			nil,
			nil,
			0,
		},
		0,
		0,
		0,
	}
	return
}