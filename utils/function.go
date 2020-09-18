package utils

import (
	"strconv"
	"strings"
	"time"
)

func IndexForOne(i int, p, limit int64) int64 {
	s := strconv.Itoa(i)
	index, _ := strconv.ParseInt(s, 10, 64)
	return (p-1)*limit + index + 1
}

func IndexAddOne(i interface{}) int64 {
	index, _ := ToInt64(i)
	return index + 1
}

func IndexDecrOne(i interface{}) int64 {
	index, _ := ToInt64(i)
	return index - 1
}

func StringReplace(str,old,new string) string {
	return  strings.Replace(str,old,new,-1)
}

func StringToTime(date interface{}) time.Time{
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	ret, _ := time.ParseInLocation(timeLayout, date.(string), loc)
	return ret
}

func TimeStampToTime (timeStamp int32) time.Time{
	return time.Unix(int64(timeStamp), 0)
}
