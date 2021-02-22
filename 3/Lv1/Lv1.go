//并发求1~50000以内的素数。注意，这里的并发是为了加速程序运行的速度（压榨多核cpu），无意义的并发不得分。
//
//加分项：
//- 使用同步机制，而不是无意义的等待
//- 使用时间复杂度不超过$O(n\log n)$的算法，并且用并发优化
package Lv1

import (
	"fmt"
	"math"
	"sync"
)

var wg sync.WaitGroup

/****************判断一个数是否是素数****************
* 输入：
*	x int 需要判断的数
* 输出：
*	是则返回true
************************************************/
func isPrime(x int) bool {
	if x<=1 {
		return false
	}
	for i:=2;i<=int(math.Sqrt(float64(x)));i++ {
		if 0==x%i {
			return false
		}
	}
	return true
}

/******以打印的方式输出某个区间(均为闭区间)的里面的所有素数*****
* 输入：
*	begin int 区间的左值
*	end	  int 区间的右值
*	wg	  sync.WaitGroup 用于支持并发
* 输出：无
****************************************************/
func take(begin,end int) {//统一认为begin为奇数，end为偶数
	defer wg.Done()
	for i:=begin;i<=end;i+=2 {//素数一定不是偶数
		if isPrime(i) {
			fmt.Printf("%05d\n",i)
		}
	}
}

func Test()  {
	for i:=1;i<=50000;i+=10000 {
		wg.Add(1)
		go take(i,i+10000)
	}
	wg.Wait()
}