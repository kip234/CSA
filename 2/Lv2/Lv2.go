//写一个接收者函数: func Receiver(v interface{}) { switch v.(Type) case ... }
//该接收者能够判断传入参数的类型，并作出不同的反应
package Lv2

import (
	"fmt"
	"reflect"
)

func Receiver(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("这个是int")
	case string:
		fmt.Println("这个是string")
	case float64:
		fmt.Println("这个是float64")
	case byte:
		fmt.Println("这个是byte")
	case int64:
		fmt.Println("这个是int64")
	case bool:
		fmt.Println("这个是bool")
	case int32:
		fmt.Println("这个是int32")
	case float32:
		fmt.Println("这个是float32")
	case complex64:
		fmt.Println("这个是complex64")
	case complex128:
		fmt.Println("这个是complex128")
	case int8:
		fmt.Println("这个是int8")
	case int16:
		fmt.Println("这个是int16")
	case uint:
		fmt.Println("这个是uint")
	case uint16:
		fmt.Println("这个是uint16")
	case uint32:
		fmt.Println("这个是uint32")
	case uint64:
		fmt.Println("这个是uint64")
	case uintptr:
		fmt.Println("这个是uintptr")
	default:
		fmt.Println("这个是复合类型或者是函数")
	}
}

//使用 内置的反射
func Receiver1(data interface{}){
	fmt.Println("这是一个",reflect.TypeOf(data))
}