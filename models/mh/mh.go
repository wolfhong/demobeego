package mh

import (
	"encoding/json"
)

// 一般返回结构
type Ret struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageRecords struct {
	Total  int         `json:"total"`
	Limit  int         `json:"limit"`
	Pageno int         `json:"pageno"`
	Data   interface{} `json:"data"`
}

func NewRet() *Ret {
	return &Ret{Msg: "ok"}
}

func (r *Ret) Error(err error) *Ret {
	r.Code = 400
	r.Msg = err.Error()
	return r
}

func (r *Ret) SetData(data interface{}) *Ret {
	r.Data = data
	return r
}

func (r *Ret) SetMsg(msg string) *Ret {
	r.Msg = msg
	return r
}

func (r *Ret) SetCode(i int) *Ret {
	r.Code = i
	return r
}

func (r *Ret) ToBytes() []byte {
	b, _ := json.Marshal(r)
	return b
}
