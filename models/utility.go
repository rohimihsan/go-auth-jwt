package models

type ResponseResult struct {
	Error  string      `json:"error"`
	Result string      `json:"result"`
	Data   interface{} `json:"data"`
}
