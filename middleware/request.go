package middleware

import (
	"github.com/labstack/echo/v4"
)

type (
	Request  struct{}
	Response struct {
		Code    int    `json:"code"`
		Status  string `json:"status"`
		Message string `json:"message"`
		Data    interface{}
	}

	PagingResponse struct {
		ListResult interface{} `json:"listresult"`
		Page       int         `json:"Page"`
		TotalData  int         `json:"totaldata"`
		TotalPage  int         `json:"totalpage"`
	}
)

func (r *Request) BindAndValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	err := c.Validate(i)
	return err
}

func (r *Request) Send(c echo.Context, code int, status, msg string, data interface{}) error {
	if data == nil {
		data = []string{}
	}

	outType := c.QueryParam("output")
	if outType == "xml" {
		return r.SendXML(c, code, status, msg, data)
	}

	return r.SendJSON(c, code, status, msg, data)
}

func (r *Request) SendJSON(c echo.Context, code int, status string, msg string, data interface{}) error {
	return c.JSON(code, map[string]interface{}{"code": code, "status": status, "message": msg, "data": data})
}

func (r *Request) SendXML(c echo.Context, code int, status, msg string, data interface{}) error {
	res := &Response{
		Code:    code,
		Status:  status,
		Message: msg,
		Data:    data,
	}
	return c.XML(code, res)
}
