package util

import (
	"strconv"
)

func PaginationString(columns []string, columnIdx int32, order string, page int32) string {
	stmt := "order by " + columns[columnIdx-1] + " " + order + " limit 20 offset " + strconv.Itoa(int((page-1)*20))
	return stmt
}

func DataFromTo(total, page int32, dataLen int) (int32, int32) {
	from := (page-1)*20 + 1
	to := (page-1)*20 + int32(dataLen)
	if from > total {
		return 0, 0
	}
	if to > total {
		to = total
	}
	return from, to
}
