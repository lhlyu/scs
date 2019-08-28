# scs
SQL条件组件

> 减少if-else语句拼接字符串

## 示例

```go
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

	sql.Reset()
	sql.AddOp("age",[]int{1,2,5,3},whr.AndIn)
	fmt.Println(sql.GetString(),sql.GetParams())     // 打印字符串  和 参数
	// 输出: select * from demo where 1=1  and age in ( ?,?,?,? ) [1 2 5 3]

}
```

## 不足，日后完善


1.  ~~in 的支持，参数数组和切片~~
2.  参数的属性化
3.  根据结构体反射自动拼接字符串
4.  update set 语句