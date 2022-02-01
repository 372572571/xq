package xq

import "encoding/json"

var Json *JsonUtil = &JsonUtil{}

type JsonUtil struct{}

func (self *JsonUtil) Marshal(object interface{}) string {
	return string(Util.MustTake(json.Marshal(object)).([]byte))
}

func (self *JsonUtil) Unmarshal(data string, object interface{}) {
	Util.Assert(json.Unmarshal([]byte(data), object))
}
