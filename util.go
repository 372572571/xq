package xq

import jwt "github.com/372572571/xq/jwtx"
type UtilType struct {
	jwt.Jwtx
}

var Util = &UtilType{}

func (self *UtilType) Assert(args ...Any) {
	if err := args[len(args)-1]; err != nil {
		panic(err)
	}
}

func (self *UtilType) MustTake(data Any, err error) Any {
	self.Assert(err)
	return data
}
