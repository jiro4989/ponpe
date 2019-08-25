package unicode

import (
	"errors"
	"fmt"
)

var (
	DiaCriticalMarks = map[rune]rune{
		'-': '\u0361', // アルファベットではないけれど特別に
		'a': '\u0363',
		'e': '\u0364',
		'i': '\u0365',
		'o': '\u0366',
		'u': '\u0367',
		'c': '\u0368',
		'd': '\u0369',
		'h': '\u036A',
		'm': '\u036B',
		'r': '\u036C',
		't': '\u036D',
		'v': '\u036E',
		'x': '\u036F',
	}
)

// ValidateDiaCriticalMark は文字列がダイアクリティカルマークに変換可能か検証す
// る。
func ValidateDiaCriticalMark(s []rune) error {
	for _, v := range s {
		if _, ok := DiaCriticalMarks[v]; !ok {
			msg := fmt.Sprintf("%sはダイアクリティカルマークではありません。", string(v))
			return errors.New(msg)
		}
	}
	return nil
}
