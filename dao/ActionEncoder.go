package dao

// DecodeResponse 解析成 code & payload
type DecodeResponse struct {
	OperationCode uint8
	EventCode     uint16
	Data          []byte
}

// EncodeData 編碼內容，包含 code
type EncodeData struct {
	OperationCode uint8
	EventCode     uint16
	Payload       interface{}
}
