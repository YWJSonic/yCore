package constant

const (
	ConnectType_None ConnectType = iota
	ConnectType_Grpc
	ConnectType_Mqueue
	ConnectType_Restful
	ConnectType_Socket
)
