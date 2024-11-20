package vender

import (
	"fmt"
	"github.com/wanggaolin/go_lib/w"
	"strings"
)

// date: 2024/11/18
// email: brach@lssin.com

func (ch *guitar) chord_scanle_verfy(query string, v Body_Chord_Scale) {
	return_name := ch.print_query(query)
	arry := w.GoString.CommaStringFormatArry(return_name)
	if len(arry) < 3 {
		arry = append(arry, "")
	}
	if len(arry) < 3 {
		arry = append(arry, "")
	}
	if len(arry) < 3 {
		arry = append(arry, "")
	}
	if err := ch.chord_scanle_answer(v, arry); err != nil {
		fmt.Println("错误,正确答案是:", strings.Join([]string{v.A, v.B, v.C}, ","), err.Error())
	} else {
		fmt.Println("回答正确")
	}
}

func (ch *guitar) chord_scanle_answer(v Body_Chord_Scale, arry []string) (err error) {
	if v.A != arry[0] {
		err = fmt.Errorf("第一个音错误，正确的是: %v", v.A)
		return
	} else if v.B != arry[1] {
		err = fmt.Errorf("第二个音错误，正确的是: %v", v.B)
		return
	} else if v.C != arry[2] {
		err = fmt.Errorf("第三个音错误，正确的是: %v", v.C)
		return
	}
	return
}
