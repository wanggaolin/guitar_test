package vender

import (
	"math/rand"
	"time"
)

// date: 2024/11/18
// email: brach@lssin.com

func (ch *guitar) show_query(args *Body_args, items Body_query) {
	if items.Type == QUERY_SCALE && args.Chord_Scale == true { // 和弦构成音
		ch.chord_scanle_verfy(items.Query, items.Answer.(Body_Chord_Scale))
	} else if items.Type == QUERY_SCALE_X && args.Chord_Scale == true { // 和弦构成音
		ch.answer_print(items.Query, items.Answer.(string))
	} else if items.Type == QUERY_LEVEL && args.Chord_Level == true { // 和弦级别
		ch.answer_print(items.Query, items.Answer.(string))
	} else if items.Type == QUERY_EMPTY && args.Chord_Empty == true { // 和弦级别
		ch.answer_print(items.Query, items.Answer.(string))
	} else if args.ShowAll == true && args.Chord_Level == false && args.Chord_Scale == false && args.Chord_Empty == false {
		if items.Type == QUERY_SCALE { // 和弦构成音
			ch.chord_scanle_verfy(items.Query, items.Answer.(Body_Chord_Scale))
		} else { // 和弦级别 + 空弦音问答 // 通用问答
			ch.answer_print(items.Query, items.Answer.(string))
		}
	}
}

// 判断问题与最斤3次是否问过
func (ch *guitar) is_repetition(x string) bool {
	l := len(ch.DATA_QUERY)
	if l < 3 {
		for _, items := range ch.DATA_QUERY {
			if items == x {
				return true
			}
		}
	} else {
		for _, items := range ch.DATA_QUERY[l-3:] {
			if items == x {
				return true
			}
		}
	}
	return false
}

func (ch *guitar) Main(args *Body_args) {
	queryData := ch.query()
	ch.NUMBER = 0
	for {
		for _, items := range queryData {
			rand.Seed(time.Now().UnixNano())
			if rand.Intn(5) == 2 {
				if ch.is_repetition(items.Query) { // 避免一个问题重复提问
					continue
				}
				ch.show_query(args, items)
			}
		}
	}
}
