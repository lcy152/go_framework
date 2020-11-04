package framework

import (
	"encoding/json"
	"net/http"
)

// Success response
func (c *Context) Success(data interface{}) {
	c.W.Header().Set("content-Type", "application/json; charset=UTF-8")
	c.W.WriteHeader(http.StatusOK)
	response := struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{200, "success", data}
	res, _ := json.Marshal(response)
	c.W.Write(res)
}

// SuccessList response
func (c *Context) SuccessList(code int, count int, data interface{}) {
	c.W.Header().Set("content-Type", "application/json; charset=UTF-8")
	c.W.WriteHeader(http.StatusOK)
	response := struct {
		Code  int         `json:"code"`
		Msg   string      `json:"msg"`
		Count int         `json:"count"`
		Data  interface{} `json:"data"`
	}{200, "success", count, data}
	res, _ := json.Marshal(response)
	c.W.Write(res)
}

// Error response
func (c *Context) Error(code int, msg string) {
	c.W.Header().Set("content-Type", "application/json; charset=UTF-8")
	c.W.WriteHeader(http.StatusNonAuthoritativeInfo)
	response := struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{code, msg}
	data, _ := json.Marshal(response)
	_, _ = c.W.Write(data)
}

// Warn response
func (c *Context) Warn(code int, msg string) {
	c.W.Header().Set("content-Type", "application/json; charset=UTF-8")
	c.W.WriteHeader(http.StatusCreated)
	response := struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{code, msg}
	data, _ := json.Marshal(response)
	_, _ = c.W.Write(data)
}
