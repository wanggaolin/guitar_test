package vender

import (
	"fmt"
)

// date: 2024/11/18
// email: brach@lssin.com

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
