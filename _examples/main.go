package main

import (
	"fmt"
	"github.com/lhlyu/scs"
	"github.com/lhlyu/scs/whr"
)

func main() {
	sql := scs.New("select * from demo where 1=1 ")
	sql.Adds("and status = 1")
	fmt.Println(sql.GetString())
	// 输出: select * from demo where 1=1 and status = 1

	sql.AddOp("name","张三",whr.AndEq)  // 相等
	fmt.Println(sql.GetString(),sql.GetParams())     // 打印字符串  和 参数
	// 输出: select * from demo where 1=1 and status = 1 and name = ?  [张三]

	sql.AddOp("age",12,whr.AndGt)       // 大于
	sql.AddOp("age",18,whr.AndLtEq)     // 小于等于
	fmt.Println(sql.GetString(),sql.GetParams())     // 打印字符串  和 参数
	// 输出: select * from demo where 1=1 and status = 1 and name = ?  and age > ?  and age <= ?  [张三 12 18]

	sql.Reset()  // 重置，回到最初New的状态
	fmt.Println(sql.GetString(),sql.GetParams())     // 打印字符串  和 参数
	// 输出: select * from demo where 1=1  []

	// 代码块  and (...)
	sql.AndBlock(func() {
		sql.AddOp("sex",1,whr.Eq)
		sql.AddOp("sex",2,whr.OrEq)
	})
	fmt.Println(sql.GetString(),sql.GetParams())     // 打印字符串  和 参数
	// 输出: select * from demo where 1=1  and ( sex = ?  or sex = ? )  [1 2]
}
