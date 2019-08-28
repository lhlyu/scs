package whr

import (
	"github.com/lhlyu/scs"
	"strings"
)

func logic(lg,column,symbol string,placeholderCount int) string{
	placeholder := scs.PLACEHOLDER
	if placeholderCount > scs.ONE{
		placeholder = strings.TrimRight(strings.Repeat("?,",placeholderCount),",")
	}
	if lg == scs.EMPTY{
		return strings.Join([]string{scs.ONE_SPACE,column,scs.ONE_SPACE,symbol,scs.ONE_SPACE,placeholder,scs.ONE_SPACE}, scs.EMPTY)
	}
	return strings.Join([]string{scs.ONE_SPACE,lg,scs.ONE_SPACE,column,scs.ONE_SPACE, symbol,scs.ONE_SPACE,placeholder,scs.ONE_SPACE}, scs.EMPTY)
}

func logicCompare(lg,column,symbol string) string{
	return logic(lg,column,symbol,scs.ONE)
}

// column = ?
func Eq(column string,placeholderCount int) string{
	return logicCompare(scs.EMPTY,column,scs.EQ)
}

// column <> ?
func NotEq(column string,placeholderCount int) string{
	return logicCompare(scs.EMPTY,column,scs.NOT_EQ)
}

// column > ?
func Gt(column string,placeholderCount int) string{
	return logicCompare(scs.EMPTY,column,scs.GT)
}

// column >= ?
func GtEq(column string,placeholderCount int) string{
	return logicCompare(scs.EMPTY,column,scs.GT_EQ)
}

// column < ?
func Lt(column string,placeholderCount int) string{
	return logicCompare(scs.EMPTY,column,scs.LT)
}

// column <= ?
func LtEq(column string,placeholderCount int) string{
	return logicCompare(scs.EMPTY,column,scs.LT_EQ)
}

// column like ?
func Like(column string,placeholderCount int) string{
	return logicCompare(scs.EMPTY,column,scs.LIKE)
}

// column not like ?
func NotLike(column string,placeholderCount int) string{
	return logicCompare(scs.EMPTY,column,scs.NOT_LIKE)
}

// column in (?,?,?...)
func In(column string,placeholderCount int) string{
	s := logic(scs.EMPTY,column,scs.IN + " (",placeholderCount)
	s += ")"
	return s
}

// column not in (?,?,?...)
func NotIn(column string,placeholderCount int) string{
	s := logic(scs.EMPTY,column,scs.NOT_IN + " (",placeholderCount)
	s += ")"
	return s
}



// and column = ?
func AndEq(column string,placeholderCount int) string{
	return logicCompare(scs.AND,column,scs.EQ)
}

// or column = ?
func OrEq(column string,placeholderCount int) string{
	return logicCompare(scs.OR,column,scs.EQ)
}

// and column <> ?
func AndNotEq(column string,placeholderCount int) string{
	return logicCompare(scs.AND,column,scs.NOT_EQ)
}

// or column <> ?
func OrNotEq(column string,placeholderCount int) string{
	return logicCompare(scs.OR,column,scs.NOT_EQ)
}

// and column > ?
func AndGt(column string,placeholderCount int) string{
	return logicCompare(scs.AND,column,scs.GT)
}

// or column > ?
func OrGt(column string,placeholderCount int) string{
	return logicCompare(scs.OR,column,scs.GT)
}

// and column >= ?
func AndGtEq(column string,placeholderCount int) string{
	return logicCompare(scs.AND,column,scs.GT_EQ)
}

// or column >= ?
func OrGtEq(column string,placeholderCount int) string{
	return logicCompare(scs.OR,column,scs.GT_EQ)
}

// and column < ?
func AndLt(column string,placeholderCount int) string{
	return logicCompare(scs.AND,column,scs.LT)
}

// or column < ?
func OrLt(column string,placeholderCount int) string{
	return logicCompare(scs.OR,column,scs.LT)
}

// and column <= ?
func AndLtEq(column string,placeholderCount int) string{
	return logicCompare(scs.AND,column,scs.LT_EQ)
}

// or column <= ?
func OrLtEq(column string,placeholderCount int) string{
	return logicCompare(scs.OR,column,scs.LT_EQ)
}

// and column like ?
func AndLike(column string,placeholderCount int) string{
	return logicCompare(scs.AND,column,scs.LIKE)
}

// and column not like ?
func AndNotLike(column string,placeholderCount int) string{
	return logicCompare(scs.AND,column,scs.NOT_LIKE)
}

// or column like ?
func OrLike(column string,placeholderCount int) string{
	return logicCompare(scs.OR,column,scs.LIKE)
}

// or column not like ?
func OrNotLike(column string,placeholderCount int) string{
	return logicCompare(scs.OR,column,scs.NOT_LIKE)
}

// and column in (?,?,?...)
func AndIn(column string,placeholderCount int) string{
	s := logic(scs.AND,column,scs.IN + " (",placeholderCount)
	s += ")"
	return s
}

// or column not in (?,?,?...)
func OrNotIn(column string,placeholderCount int) string{
	s := logic(scs.OR,column,scs.IN + " (",placeholderCount)
	s += ")"
	return s
}