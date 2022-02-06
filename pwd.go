package xq

import sbab "github.com/372572571/xq/sbab"

type Pwd struct {
}

var PwdUtil = &Pwd{}

func (u *Pwd) Encryption(mask, source string) string {
	te := sbab.New(sbab.SetMask(mask), sbab.SetSource(source))
	return te.Encryption()
}

// mask 掩 in 输入密码 source 保存密码
func (u *Pwd) Check(mask, in, sou string) bool {
	return sbab.Check(in,sou,mask)
}

func (u *Pwd) Keyer(mask  string) string {
	return sbab.Keyer(mask)
}
