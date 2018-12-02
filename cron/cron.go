/**
 * @author   : WZP
 * @createAt : 2018/12/1
 */

package cron

import "github.com/gorhill/cronexpr"

func Parse(express string)  (*cronexpr.Expression, error){
	return cronexpr.Parse(express)
}


 
 