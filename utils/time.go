package utils

import "time"

const TIME_FORMAT_ONE ="2006年01月02日 15:30:01"
const TIME_FORMAT_TWO ="2006/01/02 15:30:01"
const TIME_FORMAT_THREE ="2006.01.02 15:30:01"
const TIME_FORMAT_FOUR ="2006-01-02 15:30:01"

//时间的格式化
func TimeFormat(t int64,format string)string  {
	return time.Unix(t,0).Format(format)
}
