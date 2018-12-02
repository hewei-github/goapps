/**
 * @author   : WZP
 * @createAt : 2018/12/2
 */

package iotax

const a = iota
const b = iota

const stra string  = "12321321"
const str , strb = "12","1233"
// error
// const obj  Object = new(Object)
// var obj Object

const (
	zero = iota
	one  = iota
	two
	three
)

const (
	 z = iota * 2
	 bz
	 cz
	 dz
)

const (
	 ba = iota
	 bb = iota * 3
	 bc
	 bd
)

const (
	bx = iota
	_
	be
	bf
)

type Object struct {
	name string
}

func init()  {
	const A = a
	const B = b
	println(A,B)
	println(zero)
	println(one)
	println(two)
	println(three)
	println(z,bz,cz,dz)
}


 
 