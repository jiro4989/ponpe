package main

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/jiro4989/ponpe/unicode"
)

type (
	Config struct {
		List              bool `docopt:"-l,--list"`
		All               bool `docopt:"all"`
		DiacriticalMark   bool `docopt:"diacritical_mark,dm"`
		CyrillicAlphabets bool `docopt:"cyrillic_alphabets,ca"`
		Word              string
		Words             []string
	}
	ErrorCode int
)

const (
	doc = `ponpe prints ponpe of text.

Usage:
	ponpe [options] <word>
	ponpe [options] <word> <words>...
	ponpe [-l | --list] (all | diacritical_mark | dm | cyrillic_alphabets | ca)
	ponpe -h | --help
	ponpe -v | --version

Examples:
	ponpe ponponpain haraita-i
	ponpe ____ dddd aaaa tttt eeee
	echo ____ | ponpe date
	ponpe --list all

Options:
	-h --help       このヘルプを出力する。
	-v --version    バージョン情報を出力する。`
)

const (
	errorCodeOk ErrorCode = iota
	errorCodeFailedBinding
	errorCodeIllegalAlphabet
	errorCodeIllegalConverter
	errorCodeIllegalCommand
)

func main() {
	os.Exit(int(Main(os.Args)))
}

func Main(argv []string) ErrorCode {
	parser := &docopt.Parser{}
	args, _ := parser.ParseArgs(doc, argv[1:], Version)
	config := Config{}
	err := args.Bind(&config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return errorCodeFailedBinding
	}

	if config.List {
		return cmdList(config)
	}

	return cmdJoin(config)
}

func cmdJoin(config Config) ErrorCode {
	var word []rune
	var marks [][]rune
	for _, orgMark := range config.Words {
		// 結合文字に変換可能かチェック
		mark := []rune(orgMark)
		if err := unicode.ValidateCombindingCharacterMap(mark); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return errorCodeIllegalAlphabet
		}

		// 結合文字に変換
		mark = unicode.ToCombindingCharacterMap(mark)
		// 結合先の文字よりも、結合文字が多くてはならない
		w, m := deleteOverSize([]rune(config.Word), mark)

		// wordは常に1つのため上書きで良い
		word = w
		marks = append(marks, m)
	}

	// 結合して出力
	s := joinWords(word, marks...)
	fmt.Println(s)
	return errorCodeOk
}

func deleteOverSize(w, m []rune) ([]rune, []rune) {
	if len(w) < len(m) {
		m = m[:len(w)]
	}
	return w, m
}

func joinWords(w []rune, m ...[]rune) string {
	if len(m) < 1 {
		return string(w)
	}

	var s string
	for i := 0; i < len(w); i++ {
		s += string(w[i])
		for _, mm := range m {
			if len(mm) <= i {
				continue
			}
			s += string(mm[i])
		}
	}
	return s
}

func cmdList(config Config) ErrorCode {
	var converter map[rune]rune
	if config.All {
		converter = unicode.CombindingCharacterMap
	} else if config.DiacriticalMark {
		converter = unicode.DiaCriticalMarks
	} else if config.CyrillicAlphabets {
		converter = unicode.CyrillicAlphabets
	} else {
		// 到達しないはず
		return errorCodeIllegalConverter
	}

	for k, v := range converter {
		w, j := string(k), string(v)
		s := fmt.Sprintf("%s  %s %d u%.4x", w, j, v, v)
		fmt.Println(s)
	}
	return errorCodeOk
}
