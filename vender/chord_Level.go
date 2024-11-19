package vender

import (
	"fmt"
	"strings"
)

// date: 2024/10/30
// email: brach@lssin.com

func (ch *guitar) answer_print(query string, answer string) {
	var return_code string
	fmt.Print(ch._add_tips(query))
	fmt.Scanln(&return_code)
	if strings.ToLower(return_code) == strings.ToLower(answer) {
		fmt.Println("回答正确")
	} else {
		fmt.Println("错误,正确答案是:", answer)
	}
}

func (ch *guitar) _add_tips(query string) string {
	ch.DATA_QUERY = append(ch.DATA_QUERY, query)
	ch.NUMBER = ch.NUMBER + 1
	tips := fmt.Sprintf("第%d题: ", ch.NUMBER)
	query = tips + query
	return query
}
