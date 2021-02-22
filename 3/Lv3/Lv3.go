//使用管道，自己实现一个可以简单通知关闭的Context函数。
//
//你需要至少实现以下方法
//
//- 初始化方法（如上面的context.WithCancel）
//- 通知关闭的方法（如上面的context.Done）
//
//注意这里不是让你完全写的和context标准库一样，只要能实现通知关闭即可。
//
//记得附上使用手册。
package Lv3

import "fmt"

type Context struct {
	Parent *Context
	Alive chan struct{}
}

var EmptyContext = &Context{
	nil,
	make(chan struct{}),
}

func Add(c *Context) (*Context,func()) {//产生子节点
	if nil==c {
		fmt.Println("illegal parent")
		return nil,nil
	}
	select {
	case <-c.Over():
		fmt.Println("parent is ended")
		return nil,nil
	default:

	}
	child:=Context{
		c,
		make(chan struct{}),
	}
	return &child, func() {
		close(child.Alive)
	}
}

func (c *Context)Over() <-chan struct{}{
	select {
	case <-c.Alive://自己已经挂了
		return c.Alive
	default:
	}
	for now:=c.Parent;now!=nil;now=now.Parent {//向上查找
		select {
		case <-now.Alive://某个节点已经挂了
			close(c.Alive)//自尽
			return c.Alive
		default:
		}
	}
	return c.Alive
}