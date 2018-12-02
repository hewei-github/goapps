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
	_ "github.com/hewei-github/goapps/iotax"
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
		t  * time.Timer
		express string
	)
	// 秒 分 时 日 月 周 年 [enum timestamp]
	express = "20 * * * * * *"
	var info = "echo times :"
	now = time.Now()
	if exp,err = cron.Parse(express); nil != err{
		println(err)
		return
	}
	var j = 0
	v = exp.Next(now)
	println("触发时间: ",v.String())
	t = time.AfterFunc(v.Sub(time.Now()), func() {
		var ix = info + string(i)
		vars.Echo(ix)
		i = 1
		vars.Echo("开启定时任务了")
	})
	for 1 != i {
		time.Sleep(time.Duration(1) * time.Second)
		println("time stamp : ",v.String())
		println("time now : ",time.Now().String())
		// println("timestamp : ",v.Unix())
		j++
		println("sleep times : ", j)
	}
	t.Stop()
}
