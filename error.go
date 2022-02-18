package xq

// 错误工具
var Error *ErrorUtil = &ErrorUtil{}

type ErrorUtil struct{}

// 创建
func (self *ErrorUtil) New(template XqResp, message string) error {
	var err = new(XqResp)
	err.Code = template.Code
	err.Message = message
	return err
}
