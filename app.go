/**
 * @author   : WZP
 * @createAt : 2018/12/1
 */

package main

import (
	"github.com/gorhill/cronexpr"
	"github.com/hewei-github/goapps/cron"
	"github.com/hewei-github/goapps/vars"
	"time"
)

import (
	_ "github.com/hewei-github/goapps/cmds"
	/*_ "github.com/hewei-github/goapps/enum"*/
)

func main() {
	vars.Echo("golang hello world")
	var i = 0
	var (
		exp *cronexpr.Expression
		err error
		now time.Time
		v  time.Time
		express string
	)
	// 秒 分 时 日 月 周 年 [enum timestamp]
	express = "* 2 * * * *"
	var info = "echo times :"
	for 1 != i {
		now = time.Now()
		time.Sleep(time.Duration(1) * time.Second)
		if exp,err = cron.Parse(express); nil != err{
			println(err)
		}
		v = exp.Next(now)
		println("time stamp : ",v.String())
		println("time local : ",string(v.Local().UnixNano()))
		println("timestamp : ",string(v.Unix()))
		println("time : ", i)
		time.AfterFunc(v.Sub(time.Now()), func() {
			var ix = info + string(i)
			vars.Echo(ix)
			i = 1
		})
	}
}
