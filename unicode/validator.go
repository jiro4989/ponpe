package unicode

import (
	"errors"
	"fmt"
)

// ValidateDiaCriticalMark は文字列がダイアクリティカルマークに変換可能か検証す
// る。
func ValidateDiaCriticalMark(s []rune) error {
	return validate(s, DiaCriticalMarks, "ダイアクリティカルマーク")
}

// ValidateCyrillicAlphabets は文字列がキリル文字に変換可能か検証する。
func ValidateCyrillicAlphabets(s []rune) error {
	return validate(s, CyrillicAlphabets, "キリル文字")
}

// ValidateCombindingCharacterMap は文字列が結合文字に変換可能か検証する。
func ValidateCombindingCharacterMap(s []rune) error {
	return validate(s, CombindingCharacterMap, "結合文字")
}

func validate(s []rune, m map[rune]rune, mn string) error {
	for _, v := range s {
		if v == ' ' {
			continue
		}
		if _, ok := m[v]; !ok {
			msg := fmt.Sprintf("%sは%sではありません。", string(v), mn)
			return errors.New(msg)
		}
	}
	return nil
}
