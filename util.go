package xq

import (
	jwt "github.com/372572571/xq/jwtx"
	"github.com/372572571/xq/lang"
)

type UtilType struct {
	jwt.Jwtx
}

var Util = &UtilType{}

func (self *UtilType) Assert(args ...lang.Any) {
	if err := args[len(args)-1]; err != nil {
		panic(err)
	}
}

func (self *UtilType) MustTake(data lang.Any, err error) lang.Any {
	self.Assert(err)
	return data
}
