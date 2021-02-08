//该计算器可以执行两个数字和一个计算符号的计算。
package Task1

import "fmt"

// @title function
// @description 直接实现目标功能
// @auth 刘帅
// @param 无
// @return 无
func Task()  {
	var (
		a int32
		b byte//运算符
		c int32
	)
	fmt.Scanf("%d%c%d",&a,&b,&c)//开始输入
	switch b {//判断运算符
	case '+':
		fmt.Println(a+c)
	case '-':
		fmt.Println(a-c)
	case '*':
		fmt.Println(a*c)
	case '/':
		fmt.Println(a/c)
	default:
		fmt.Printf("Illeagle operation \"%s\" !\n",string(b))
	}
}