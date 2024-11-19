package vender

// date: 2024/10/30
// email: brach@lssin.com

const QUERY_LEVEL = "CHORD_LEVEL" // 和弦级别问答
const QUERY_SCALE = "CHORD_SCALE" // 和弦级构成音
const QUERY_EMPTY = "CHORD_EMPTY" // 空弦音

// 和弦级别
type Body_Chord_Level struct {
	Level int64  // 级别
	Chord string // 和弦
}

// 和弦音阶组成
type Body_Chord_Scale struct {
	A string
	B string
	C string
}

type guitar struct {
	DATA_CHORD_LEVEL map[string][]Body_Chord_Level //和弦级数数据
	DATA_CHORD_SCALE map[string]Body_Chord_Scale   //和弦音阶组成
	DATA_CHORD_EMPTY map[string]string             //空弦音
	DATA_QUERY       string                        //上一个问题
}

type Body_args struct {
	Version     bool
	Chord_Level bool // 和弦级数问答
	Chord_Scale bool // 和弦级够成音问答
	Chord_Empty bool // 吉他空弦音
	Query       bool // 生成在线问答题
	ShowAll     bool // 题库随机问答
}

type Body_query struct {
	Query  string      // 问题
	Type   string      // 问题类型
	Answer interface{} // 答案
}
