package counttool

import (
	"context"
	"time"
)

// 時間內計數工具
type CounePreDuration struct {
	time     time.Duration
	incrTime []int64
	msgChan  chan int
	resChan  chan int
	ctx      context.Context
	stop     context.CancelFunc
}

func NewCounePreDuration(d time.Duration) *CounePreDuration {
	if d < time.Second {
		d = time.Second
	}

	ctx, stop := context.WithCancel(context.TODO())
	obj := &CounePreDuration{
		time:    d,
		msgChan: make(chan int),
		resChan: make(chan int),
		ctx:     ctx,
		stop:    stop,
	}
	go obj.handle()
	return obj
}

func (self *CounePreDuration) handle() {
	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()
	for {
		select {
		case <-self.ctx.Done():
			return
		case opcode := <-self.msgChan: // req msg
			switch opcode {
			case 0: // 未定義
			case 1: // incr
				self.incrTime = append(self.incrTime, time.Now().UnixMilli())
			case 2:
				self.resChan <- len(self.incrTime)
			}
		case <-ticker.C: // update
			if len(self.incrTime) > 0 {
				now := time.Now().UnixMilli()
				if time.Duration(now-self.incrTime[0])*time.Millisecond >= self.time {
					for idx, incrTime := range self.incrTime {
						if time.Duration(now-incrTime)*time.Millisecond < self.time {
							if idx != 0 {
								self.incrTime = self.incrTime[idx:]
								break
							}
						} else if idx == len(self.incrTime)-1 {
							self.incrTime = self.incrTime[idx:]
						}
					}
				}
			}
		}
	}
}

// 停止內部處理結束時務必執行
func (self *CounePreDuration) Stop() {
	self.stop()
}

// 新增計數
func (self *CounePreDuration) Incr() {
	self.msgChan <- 1
}

// 當前計數
func (self *CounePreDuration) GetCount() int {
	self.msgChan <- 2
	return <-self.resChan
}
