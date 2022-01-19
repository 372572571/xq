package test

import (
	"encoding/hex"
	"fmt"
	"testing"
	"xq/encryption/sbab"
	"xq/internal"
)

func est_sbab(t *testing.T) {
	var key = string(internal.Util.MustTake(hex.DecodeString("6368616e67652074686973e070e17323")).([]byte))
	var t_s = sbab.New(sbab.SetSource("lul"), sbab.SetMask(key))
	var t_e = t_s.Encryption()
	if sbab.Check("lyl", t_e, key) {
		fmt.Println("ok")
	} else {
		fmt.Println("error")
	}
}
