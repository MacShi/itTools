package utils

import (
	"sync"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
var (
	r          *Result
	once       sync.Once
	ResultUtil = getResult()
)
func getResult() *Result {
	// 如果不使用 sync.Once，而是每次调用 getResult 都创建一个新的 Result 对象，那么可能会造成不必要的资源浪费，因为实际上只需要一个全局的 Result 对象。
	once.Do(func() {
		r = &Result{}
	})
	return r
}
//Success 方法定义，属于 *Result 类型的方法集合
func (r *Result) Success(data interface{}) *Result {
	r.Code = 0
	r.Msg = "success"
	r.Data = data
	return r
}
//Failure 方法定义，属于 *Result 类型的方法集合
func (r *Result) Failure(msg string) *Result {
	r.Code = -1
	r.Msg = msg
	r.Data = nil
	return r
}
