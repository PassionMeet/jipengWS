package protocol

type MessageBase struct {
	RmtAddr string      `json:"rmt_addr"`
	Type    string      `json:"type"`
	Msg     interface{} `json:"msg"`
}

type MessageImpl_ struct{}
