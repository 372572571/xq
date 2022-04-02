package xq

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/372572571/commpb"
	"github.com/372572571/xq/lang"
)

type XqResp = commpb.Status

// 默认错误
type DefResp struct {
	OK       XqResp `success:"0000001"`
	Unknown  XqResp `error:"0001001"`
	Param    XqResp `error:"0001002"`
	Network  XqResp `error:"0001003"`
	Server   XqResp `error:"0001004"`
	Token    XqResp `error:"0001005"`
	NotFound XqResp `error:"0001006"`
}

type XqRespUtil struct {
	Def *DefResp
}

var RespUtil = newRespUtil()

func newRespUtil() *XqRespUtil {
	var resp = &XqRespUtil{Def: new(DefResp)}
	resp.SetRespCode(resp.Def)
	return resp
}

// 设置返回码
func (self *XqRespUtil) SetRespCode(group lang.Any) {
	var inject func(re, rs *regexp.Regexp, group lang.Any)

	inject = func(re, rs *regexp.Regexp, group lang.Any) {
		elem := reflect.Indirect(reflect.ValueOf(group))
		data := reflect.TypeOf(elem.Interface())

		for i := 0; i < data.NumField(); i++ {
			info := data.Field(i)
			if info.Anonymous {
				addr := elem.FieldByName(info.Name).Addr()
				inject(re, rs, addr.Interface())
				continue
			}

			code := string(re.Find([]byte(info.Tag)))
			mess := string(rs.Find([]byte(info.Tag)))
			resp := XqResp{
				Code:    int32(Util.MustTake(strconv.Atoi(code)).(int)),
				Message: strings.ToUpper(mess + " " + info.Name),
				Details: []string{},
			}
			pair := reflect.ValueOf(resp)
			elem.FieldByName(info.Name).Set(pair)
		}
	}

	re := Util.MustTake(regexp.Compile("[[:digit:]]+")).(*regexp.Regexp)
	rs := Util.MustTake(regexp.Compile("[a-z]+")).(*regexp.Regexp)

	for _, v := range []lang.Any{group} {
		if v == nil {
			continue
		}
		inject(re, rs, v)
	}
}

// 错误信息打包
func (u *XqRespUtil) Pkg(e XqResp, details ...string) XqResp {
	var err = new(XqResp)
	err.Code = e.Code
	err.Message = e.Message
	err.Details = []string{}

	if len(details) > 0 {
		err.Details = append(err.Details, details...)
	}

	return *err
}

func (u *XqRespUtil) PkgToString(e XqResp, details ...string) string {
	var err = new(XqResp)
	err.Code = e.Code
	err.Message = e.Message
	err.Details = []string{}

	if len(details) > 0 {
		err.Details = append(err.Details, details...)
	}
	return Json.Marshal(err)
}
