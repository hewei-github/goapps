/**
 * @author   : WZP
 * @createAt : 2018/12/1
 */

package vars

import (
	"github.com/hewei-github/goapps/enum"
	"reflect"
	"unsafe"
)

// 常量
const NAME = "golang"
const NUM  = 1

// 函数
func echo(str string)  {
	println(string(str))
}

// 自定义 类型
type Object struct {
	name string
	__attrs__ []string
}

const __ENUM__ = enum.ENUM_PACK

const FALSE  = false
const False  = false
const TRUE   = true
const True   = true
const Nil    = 0x000


func globalTest()  {
	println(NUM)
	println(NAME)
	println(echo)
	println(__ENUM__)
	// 创建自定义类型
	var obj =new(Object)
	println(obj)
	obj.name = "golang"
	println(obj)

	println(false)
	println(FALSE)
	println(False)
	println(true)
	println(TRUE)
	println(True)
	println(Nil)
	println("Nil 类型 :",reflect.TypeOf(Nil))
	println("echo 类型 :",reflect.TypeOf(echo))
	println("obj 类型 :",reflect.TypeOf(obj))
	println("大小 :",unsafe.Sizeof(echo))
}

func Echo(str string)  {
	 echo(str)
}


 
 