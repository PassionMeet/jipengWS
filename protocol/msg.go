package protocol

type MessageBase struct {
	UserID string      `json:"user_id"`
	Type   string      `json:"type"`
	Msg    interface{} `json:"msg"`
}

type MessageImpl_ struct{}
