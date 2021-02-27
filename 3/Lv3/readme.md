## 思路



## 自定义类型：

```go
type Context struct {
	Parent *Context		//指向上一级
	Alive chan struct{}	//
}
```

## 变量

```go
var EmptyContext = &Context{
	nil,//默认的根节点
	make(chan struct{}),
}
```

## 函数

| 函数原型                                | 函数含义                        |
| --------------------------------------- | ------------------------------- |
| func Add(c *Context) (*Context,func())  | 产生子节点及通知关闭的函数      |
| func (c *Context)Over() <-chan struct{} | 返回一个channel以表示当前的状态 |

## 示例

```go
package main

import (
	"fmt"
	"./Lv3"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx,cancel:=Lv3.Add(Lv3.EmptyContext)//产生下一级节点并作为参数传入下一级协程

	wg.Add(1)
	go func(c *Lv3.Context) {
		defer wg.Done()
		for{
			select {
			case <-c.Over()://如果通知结束就会取到0或空
				fmt.Println("over")
				return
			default:
				fmt.Println("working...")
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	time.Sleep(time.Second*2)
	cancel()//调用返回的函数表示结束所有子协程
	wg.Wait()
}
```

## 奇怪的知识回来了

channel，一个奇怪的东西。

```go
package main

import "fmt"

func main() {
	channel:=make(chan int)

	fmt.Println(<-channel)
}
```

> 然后就会报错，产生死锁

```
fatal error: all goroutines are asleep - deadlock!
```

然鹅


```go
package main

import "fmt"

func main() {
	channel:=make(chan int)
	close(channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
```

> 这样就不会

似乎只要是能够传值的channel，在close以后就可以不限次数地读取0值或空值(视其类型而定)
