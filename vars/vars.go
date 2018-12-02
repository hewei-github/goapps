/**
 * @author   : WZP
 * @createAt : 2018/12/1
 */

package vars

// 包引入时初始调用函数
func init()  {
	echo("初始化成功")
}

func varsTest() {
	var strMain = '和'          // Unicode char
	var strStr = "hello world" // string
	var strVar string

	println(strMain)
	println(strStr)
	strVar = "string"
	println(strVar)
}
