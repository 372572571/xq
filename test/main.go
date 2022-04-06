package main

import (
	"bytes"
	"fmt"

	"github.com/372572571/xq/jsonx"
)

// func main() {
// 	t := &time.Time{}
// 	t2 := &time.Time{}
// 	fmt.Println(t)
// 	fmt.Printf("%p \n",t)
// 	fmt.Printf("%p \n",t2)
// 	fmt.Println((*time.Time)(t))
// }

// func main() {
// 	data := make(chan bool, 2)
// 	data <- bool(true)
// 	fmt.Println(1)
// 	data <- bool(true)
// 	fmt.Println(2)
// 	go func() {
// 		t := time.NewTicker(time.Second * 10)
// 		<-t.C
// 		<-data
// 		fmt.Println("time over")
// 	}()
// 	data <- bool(true)
// 	fmt.Println(3)
// 	fmt.Println("over")
// }

type data struct {
	Value int64
}

func main() {
	// js := "{value:123}"
	var resp = bytes.NewBufferString("")
	v := &data{Value: 1235}
	jsonx.Marshal(&v, resp)
	fmt.Println(resp)
}
