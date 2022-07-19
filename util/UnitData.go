package util

import (
	"time"

	"github.com/rs/xid"
)

func UnitStr() string {
	return xid.New().String()
}

func ServerTimeNow() time.Time {
	return time.Now().UTC()
}
