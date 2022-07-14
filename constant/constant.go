package constant

const (
	ConnectType_None ConnectType = iota
	ConnectType_Grpc
	ConnectType_Mqueue
	ConnectType_Restful
	ConnectType_Socket
)

const (
	MaxUint16 = 1<<16 - 1
	MaxUint64 = 1<<64 - 1
	MaxUint   = ^uint(0)
	MinUint   = 0
	MaxInt    = int(MaxUint >> 1)
	MinInt    = -MaxInt - 1
)
