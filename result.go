package jpush

// Result is jpush response data struct
type Result struct {
	SendNo string                 `json:"sendno"`
	MsgID  string                 `json:"msg_id"`
	Error  map[string]interface{} `json:"error"`
}
