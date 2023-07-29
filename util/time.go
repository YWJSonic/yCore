package util

import (
	"time"
)

// 取得當前 server 時間 *時區為 UTC+0
func ServerTimeNow() time.Time {
	return time.Now().UTC()
}

// 計算當月天數
//
// @params time.Month 月份
//
// @params int 年分
//
// @return int 天數
func DaysIn(m time.Month, y int) int {
	return time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// 計算天數, 只關注日期不會在意時間
//
// @params time.Time 開始時間
//
// @params time.Time 結束時間
//
// @return int 天數
func CountTransDate(startTime, endTime time.Time) int {
	startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, time.UTC)
	endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 0, 0, 0, 0, time.UTC)
	return int(endTime.Sub(startTime).Hours() / 24)
}
