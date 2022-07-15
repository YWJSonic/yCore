package util

import "github.com/rs/xid"

func UnitStr() string {
	return xid.New().String()
}
