package core

type Responses interface {
	SetCode(code int)
	SetTraceID(string)
	SetMsg(string)
	SetData(interface{})
	Clone() Responses
}

type Response struct {
	// 数据集
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

type response struct {
	Response
	Data interface{} `json:"data"`
}

func (e *response) SetData(data interface{}) {
	e.Data = data
}

func (e response) Clone() Responses {
	return &e
}

func (e *response) SetTraceID(id string) {
	e.RequestId = id
}

func (e *response) SetMsg(s string) {
	e.Message = s
}

func (e *response) SetCode(code int) {
	e.Code = code
}
