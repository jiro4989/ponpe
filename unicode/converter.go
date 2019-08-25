package unicode

var (
	CombindingCharacterMap = make(map[rune]rune)
)

func init() {
	// 各アルファベットに対応する結合文字を1つのマップにまとめる
	for _, c := range []map[rune]rune{DiaCriticalMarks, DiaCriticalMarksSupplement, CyrillicAlphabets, EtcLetters} {
		for k, v := range c {
			CombindingCharacterMap[k] = v
		}
	}
}

// ToDiaCriticalMark はダイアクリティカルマークへ変換する。
func ToDiaCriticalMark(s []rune) []rune {
	return convert(s, DiaCriticalMarks)
}

// ToCyrillicAlphabets はダイアクリティカルマークへ変換する。
func ToCyrillicAlphabets(s []rune) []rune {
	return convert(s, CyrillicAlphabets)
}

// ToCombindingCharacterMap はアルファベットを結合文字へ変換する。
func ToCombindingCharacterMap(s []rune) []rune {
	return convert(s, CombindingCharacterMap)
}

func convert(s []rune, m map[rune]rune) []rune {
	var marks []rune
	for _, v := range s {
		mark := m[v]
		marks = append(marks, mark)
	}
	return marks
}
