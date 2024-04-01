package qntp

import (
	"github.com/beevik/ntp"
	"time"
)

var clockOffset time.Duration //时钟偏差

// SyncTime 同步时间 向指定ntp服务器同步时间
func SyncTime(ntpHost string) error {
	r, err := ntp.Query(ntpHost)
	if err != nil {
		return err
	}

	if err = r.Validate(); err != nil {
		return err
	}
	clockOffset = r.ClockOffset
	return nil
}

// AdjustedTime 已校正的时间
func AdjustedTime() time.Time {
	return time.Now().Add(clockOffset)
}
