package models

type Response struct {
	ErrorText string      `json:"error_text"`
	HasError  bool        `json:"has_error"`
	Resp      interface{} `json:"resp"`
}
