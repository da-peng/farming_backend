package controllers

//ErrorController 自定义错误
type ErrorController struct {
	BaseController
}

// Error404 404错误
func (c *ErrorController) Error404() {
	c.ResponseJSON(20404, "请求异常")
}

// Error501 501 错误
func (c *ErrorController) Error501() {

	c.ResponseJSON(20501, "服务端异常")
}

// ErrorDb 数据库错误
func (c *ErrorController) ErrorDb() {
	c.ResponseJSON(20505, "数据库异常")
}
