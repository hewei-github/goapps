/**
 * @author   : WZP
 * @createAt : 2018/12/1
 */

package cmds

import "os/exec"

func init()  {
	go server()
}

func server()  {
	var cmd = exec.Command("php","-S","0.0.0.0:8080")
	var (
		a []byte
		err error
	)
	if  a,err = cmd.CombinedOutput(); nil != err {
		println(err)
		panic(err)
	}
	println(string(a))
}


 
 