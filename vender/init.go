package vender

// date: 2024/10/30
// email: brach@lssin.com

var Guitar *guitar // 音阶

func init() {
	Guitar = &guitar{}
	_init_chord_level() // 初始化和弦级数
	_init_chord_scale() // 初始化和弦音阶组成
	_init_chord_empty() // 初始化空弦音
}

// 初始化和弦音阶
func _init_chord_scale() {
	Guitar.DATA_CHORD_SCALE = make(map[string]Body_Chord_Scale)
	Guitar.DATA_CHORD_SCALE["C"] = Body_Chord_Scale{"1", "3", "5"}
	Guitar.DATA_CHORD_SCALE["D"] = Body_Chord_Scale{"2", "#4", "6"}
	Guitar.DATA_CHORD_SCALE["DM"] = Body_Chord_Scale{"2", "4", "6"}
	Guitar.DATA_CHORD_SCALE["E"] = Body_Chord_Scale{"3", "#5", "7"}
	Guitar.DATA_CHORD_SCALE["EM"] = Body_Chord_Scale{"3", "7", ""}
	Guitar.DATA_CHORD_SCALE["G"] = Body_Chord_Scale{"5", "7", "2"}
	Guitar.DATA_CHORD_SCALE["F"] = Body_Chord_Scale{"4", "6", "1"}
	Guitar.DATA_CHORD_SCALE["A"] = Body_Chord_Scale{"6", "#1", "3"}
	Guitar.DATA_CHORD_SCALE["AM"] = Body_Chord_Scale{"6", "1", "3"}
}

// 初始化和弦级数
func _init_chord_level() {
	Guitar.DATA_CHORD_LEVEL = map[string][]Body_Chord_Level{}
	var chord_c []Body_Chord_Level
	var chord_g []Body_Chord_Level
	chord_c = append(chord_c, Body_Chord_Level{Level: 1, Chord: "C"})
	chord_c = append(chord_c, Body_Chord_Level{Level: 2, Chord: "DM"})
	chord_c = append(chord_c, Body_Chord_Level{Level: 3, Chord: "EM"})
	chord_c = append(chord_c, Body_Chord_Level{Level: 4, Chord: "F"})
	chord_c = append(chord_c, Body_Chord_Level{Level: 5, Chord: "G"})
	chord_c = append(chord_c, Body_Chord_Level{Level: 6, Chord: "AM"})

	chord_g = append(chord_g, Body_Chord_Level{Level: 1, Chord: "G"})
	chord_g = append(chord_g, Body_Chord_Level{Level: 2, Chord: "AM"})
	chord_g = append(chord_g, Body_Chord_Level{Level: 3, Chord: "BM"})
	chord_g = append(chord_g, Body_Chord_Level{Level: 4, Chord: "C"})
	chord_g = append(chord_g, Body_Chord_Level{Level: 5, Chord: "D"})
	chord_g = append(chord_g, Body_Chord_Level{Level: 6, Chord: "EM"})

	Guitar.DATA_CHORD_LEVEL["C"] = chord_c
	Guitar.DATA_CHORD_LEVEL["G"] = chord_g
}

// 初始化空弦音
func _init_chord_empty() {
	Guitar.DATA_CHORD_EMPTY = make(map[string]string)
	Guitar.DATA_CHORD_EMPTY["6"] = "E"
	Guitar.DATA_CHORD_EMPTY["5"] = "A"
	Guitar.DATA_CHORD_EMPTY["4"] = "D"
	Guitar.DATA_CHORD_EMPTY["3"] = "G"
	Guitar.DATA_CHORD_EMPTY["2"] = "B"
	Guitar.DATA_CHORD_EMPTY["1"] = "E"
}
