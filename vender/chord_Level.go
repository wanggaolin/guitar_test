package vender

import (
	"fmt"
	"strings"
)

// date: 2024/10/30
// email: brach@lssin.com

func (ch *guitar) answer_print(query string, answer string) {
	var return_code string
	fmt.Print(query)
	fmt.Scanln(&return_code)
	if strings.ToLower(return_code) == strings.ToLower(answer) {
		fmt.Println("回答正确")
	} else {
		fmt.Println("错误,正确答案是:", answer)
	}
}
