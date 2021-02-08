//请参照上面的例子，声明鸽子，复读机，柠檬精，真香怪四种接口，然后让人类类型分别实现四种
//接口。
package Lv1

import "fmt"

// person 人类 
type Person struct {
	name string //姓名 
	age int // 年龄 
	gender string // 性别
}

// 鸽子
type dove interface {
	Gugugu() // 鸽
}

func (p Person)Gugugu()  {
	fmt.Println("gu一下")
}

// 复读机
type repeater interface {
	Repeat(string) // 复读
}
func (p Person)Repeat(string2 string)  {
	fmt.Println(string2)
}

// 柠檬精
type lemonElf interface {
	Suan()//太酸了
}
func (p Person)Suan() {
	fmt.Println("好酸")
}

// 真香怪
type zhenXiang interface {
	daLian()
}
func (p Person)daLian()  {
	fmt.Println("五金打脸 (￣ε(#￣)☆╰╮(￣▽￣///)")
}