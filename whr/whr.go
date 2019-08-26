package whr

import "strings"

func logic(lg,column,symbol string,placeholderCount int) string{
	placeholder := "?"
	if placeholderCount > 1{
		placeholder = strings.TrimRight(strings.Repeat("?,",placeholderCount),",")
	}
	if lg == ""{
		return strings.Join([]string{"",column, symbol,placeholder,""}, " ")
	}
	return strings.Join([]string{"",lg,column, symbol,placeholder,""}, " ")
}

func logicCompare(lg,column,symbol string) string{
	return logic(lg,column,symbol,1)
}

// column = ?
func Eq(column string,placeholderCount int) string{
	return logicCompare("",column,"=")
}

// column <> ?
func NotEq(column string,placeholderCount int) string{
	return logicCompare("",column,"<>")
}

// column > ?
func Gt(column string,placeholderCount int) string{
	return logicCompare("",column,">")
}

// column >= ?
func GtEq(column string,placeholderCount int) string{
	return logicCompare("",column,">=")
}

// column < ?
func Lt(column string,placeholderCount int) string{
	return logicCompare("",column,"<")
}

// column <= ?
func LtEq(column string,placeholderCount int) string{
	return logicCompare("",column,"<=")
}

// column like ?
func Like(column string,placeholderCount int) string{
	return logicCompare("",column,"like")
}

// column not like ?
func NotLike(column string,placeholderCount int) string{
	return logicCompare("",column,"not like")
}

// column in (?,?,?...)
func In(column string,placeholderCount int) string{
	return logic("",column,"in",placeholderCount)
}

// column not in (?,?,?...)
func NotIn(column string,placeholderCount int) string{
	return logic("",column,"not in",placeholderCount)
}



// and column = ?
func AndEq(column string,placeholderCount int) string{
	return logicCompare("and",column,"=")
}

// or column = ?
func OrEq(column string,placeholderCount int) string{
	return logicCompare("or",column,"=")
}

// and column <> ?
func AndNotEq(column string,placeholderCount int) string{
	return logicCompare("and",column,"<>")
}

// or column <> ?
func OrNotEq(column string,placeholderCount int) string{
	return logicCompare("or",column,"<>")
}

// and column > ?
func AndGt(column string,placeholderCount int) string{
	return logicCompare("and",column,">")
}

// or column > ?
func OrGt(column string,placeholderCount int) string{
	return logicCompare("or",column,">")
}

// and column >= ?
func AndGtEq(column string,placeholderCount int) string{
	return logicCompare("and",column,">=")
}

// or column >= ?
func OrGtEq(column string,placeholderCount int) string{
	return logicCompare("or",column,">=")
}

// and column < ?
func AndLt(column string,placeholderCount int) string{
	return logicCompare("and",column,"<")
}

// or column < ?
func OrLt(column string,placeholderCount int) string{
	return logicCompare("or",column,"<")
}

// and column <= ?
func AndLtEq(column string,placeholderCount int) string{
	return logicCompare("and",column,"<=")
}

// or column <= ?
func OrLtEq(column string,placeholderCount int) string{
	return logicCompare("or",column,"<=")
}

// and column like ?
func AndLike(column string,placeholderCount int) string{
	return logicCompare("and",column,"like")
}

// and column not like ?
func AndNotLike(column string,placeholderCount int) string{
	return logicCompare("and",column,"not like")
}

// or column like ?
func OrLike(column string,placeholderCount int) string{
	return logicCompare("or",column,"like")
}

// or column not like ?
func OrNotLike(column string,placeholderCount int) string{
	return logicCompare("or",column,"not like")
}

// and column in (?,?,?...)
func AndIn(column string,placeholderCount int) string{
	return logic("and",column,"in",placeholderCount)
}

// or column not in (?,?,?...)
func OrNotIn(column string,placeholderCount int) string{
	return logic("or",column,"not in",placeholderCount)
}