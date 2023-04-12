package timecount

import (
	"time"

	"github.com/YWJSonic/ycore/util"
)

// 生成新計時器
//
// @param nowTime 當下時間
// * 寫入不會做任何調整包含時區
func NewClockTime(nowTime time.Time) ClockTime {
	return ClockTime{
		dateTime: nowTime,
	}
}

type ClockTime struct {
	tickCount uint64        //timer 計數次數
	timer     time.Duration // 時間計數
	dateTime  time.Time     // 計數當下時間
}

func (t *ClockTime) MarshalJSON() ([]byte, error) {
	alias := struct {
		Timer     time.Duration `json:"timer"`
		DateTime  time.Time     `json:"dateTime"`
		TickCount uint64        `json:"tickCount"`
	}{
		Timer:    t.timer,
		DateTime: t.dateTime,
	}

	data, err := util.Marshal(alias)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t *ClockTime) UnmarshalJSON(b []byte) error {
	alias := struct {
		Timer     time.Duration `json:"timer"`
		DateTime  time.Time     `json:"dateTime"`
		TickCount uint64        `json:"tickCount"`
	}{}

	if err := util.Unmarshal(b, &alias); err != nil {
		return err
	}

	t.timer = alias.Timer
	t.dateTime = alias.DateTime
	t.tickCount = alias.TickCount

	return nil
}

// 取得當前時間
// 根據初始化給予的 nowTime 自行計數時候的時間
//
// @return time.Time 當下時間
func (t *ClockTime) GetNowTime() time.Time {
	return t.dateTime
}

// 取得當下計數時間
//
// @return time.Duration 計數時間
func (t *ClockTime) GetTimer() time.Duration {
	return t.timer
}

// 取得當下計數時間
//
// @return time.Duration 計數時間
func (t *ClockTime) GetTickCount() uint64 {
	return t.tickCount
}

// 設定新的當前時間
//
// @param nowTime 當下時間
// * 寫入不會做任何調整包含時區
func (t *ClockTime) SetNowTime(nowTime time.Time) {
	t.dateTime = nowTime
}

// 重製計數器
func (t *ClockTime) ResetTimer() {
	t.timer = 0
	t.tickCount = 0
}

// 計數一段時間
//
// @param tickTime 要前進的時長
func (t *ClockTime) Ticket(tickTime time.Duration) {
	t.tickCount++
	t.timer += tickTime
	t.dateTime = t.dateTime.Add(tickTime)
}
