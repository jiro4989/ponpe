package main

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/jiro4989/ponpe/unicode"
)

type (
	Config struct {
		CmdJoin           bool `docopt:"join,j"`
		CmdList           bool `docopt:"list,l"`
		All               bool `docopt:"all,a"`
		DiaCriticalMark   bool `docopt:"diacritical_mark,dm"`
		CyrillicAlphabets bool `docopt:"cyrillic_alphabets,ca"`
		Word              string
		Words             []string
	}
	ErrorCode int
)

const (
	doc = `ponpe prints ponpe of text.

Usage:
	ponpe (join | j) <word> <words>...
	ponpe (list | l) (all | a | diacritical_mark | dm | cyrillic_alphabets | ca)
	ponpe -h | --help
	ponpe -v | --version

Options:
	-h --help                     Show this screen.
	-v --version                  Show version.`
)

const (
	errorCodeOk ErrorCode = iota
	errorCodeFailedBinding
	errorCodeIllegalAlphabet
	errorCodeIllegalConverter
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

	if config.CmdJoin {
		return cmdJoin(config)
	}

	if config.CmdList {
		return cmdList(config)
	}

	return 99
}

func cmdJoin(config Config) ErrorCode {
	var word []rune
	var marks [][]rune
	for _, orgMark := range config.Words {
		// ダイアクリティカルマークに変換可能かチェック
		mark := []rune(orgMark)
		if err := unicode.ValidateDiaCriticalMark(mark); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return errorCodeIllegalAlphabet
		}

		// ダイアクリティカルマークに変換
		mark = unicode.ToDiaCriticalMark(mark)
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
	} else if config.DiaCriticalMark {
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
