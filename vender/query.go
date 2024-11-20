package vender

import (
	"bufio"
	"fmt"
	"github.com/wanggaolin/go_lib/w"
	"golang.org/x/term"
	"os"
	"runtime"
	"strings"
)

// date: 2024/11/18
// email: brach@lssin.com

func (ch *guitar) __print_query_linux() (answer string) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		w.Log.Log_error("mod error:", err.Error())
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	term := term.NewTerminal(os.Stdin, "")
	input, err := term.ReadLine()
	if err != nil {
		if err.Error() == "EOF" {
			w.ExitSuccess()
		}
		w.Log.Log_error("read error:", err.Error())
		return
	}
	answer = input
	return
}

func (ch *guitar) print_query(query string) (answer string) {
	ch.DATA_QUERY = append(ch.DATA_QUERY, query)
	ch.NUMBER = ch.NUMBER + 1
	tips := fmt.Sprintf("第%d题: ", ch.NUMBER)
	query = tips + query
	fmt.Print(query)
	if runtime.GOOS == "windows" {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		answer = input[:len(input)-1]
	} else {
		answer = ch.__print_query_linux()
	}
	return strings.TrimSpace(answer)
}

// 生成在线问题
func (ch *guitar) query() (queryData []Body_query) {
	// 和弦级数
	for melody, items := range ch.DATA_CHORD_LEVEL {
		for _, items_chord := range items {
			queryData = append(queryData, Body_query{
				Query:  fmt.Sprintf("请输入 %s调 %v和弦是_____级和弦: ", melody, items_chord.Chord),
				Answer: fmt.Sprintf("%d", items_chord.Level),
				Type:   QUERY_LEVEL,
			})

			queryData = append(queryData, Body_query{
				Query:  fmt.Sprintf("请输入 %s调 %v级和弦是______: ", melody, items_chord.Level),
				Answer: items_chord.Chord,
				Type:   QUERY_LEVEL,
			})
		}
	}
	// 和弦构成音
	for k, v := range ch.DATA_CHORD_SCALE {
		queryData = append(queryData, Body_query{
			Query:  fmt.Sprintf("%s 和弦构成音(多个音之间逗号间隔)______: ", k),
			Answer: v,
			Type:   QUERY_SCALE,
		})
		queryData = append(queryData, Body_query{
			Query:  fmt.Sprintf("%s 组成了______和弦: ", strings.Trim(v.A+","+v.B+","+v.C, ",")),
			Answer: k,
			Type:   QUERY_SCALE_X,
		})
	}

	// 和弦构空弦音
	for k, v := range ch.DATA_CHORD_EMPTY {
		queryData = append(queryData, Body_query{
			Query:  fmt.Sprintf("第 %s 弦空弦音是: ", k),
			Answer: v,
			Type:   QUERY_EMPTY,
		})
	}
	return
}
