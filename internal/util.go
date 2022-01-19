package internal

type UtilType struct{}

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
