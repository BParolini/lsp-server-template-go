package messages

type Result interface {
	string | int | bool | []any | any
}

type JsonMessage struct {
	JsonRPC string `json:"jsonrpc"`
}

type RequestMessage struct {
	JsonMessage
	Id     uint     `json:"id"`
	Method string   `json:"method"`
	Params []string `json:"params"`
}

type ResponseMessage struct {
	JsonMessage
	Id     uint          `json:"id"`
	Result Result        `json:"result"`
	Error  ResponseError `json:"error"`
}

type ResponseError struct {
	JsonMessage
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Result `json:"data"`
}

type NotificatonMessage struct {
	JsonMessage
	Method string `json:"method"`
	Params any    `json:"params"`
}

type ParamId interface {
	string | int
}

type CancelParams struct {
	Id string `json:"id"`
}
