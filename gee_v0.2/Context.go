package gee_v0_2

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

// 获取Url数据/表单数据
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key) // url和表单相同key，优先返回表单中的
}

// 设置头信息
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// 查询Url参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 设置状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) Json(statusCode int, data any) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(statusCode)

	encoder := json.NewEncoder(c.Writer)
	err := encoder.Encode(data)
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}

	// Marshal: 只支持短数据，不建议
	/*bytes, err := json.Marshal(data)
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}*/
}

// 传递string，便于提供给用户格式化
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
