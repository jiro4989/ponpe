package main

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/jiro4989/ponpe/unicode"
)

type Config struct {
	Word  string
	Words []string
}

const (
	doc = `ponpe prints ponpe of text.

Usage:
	ponpe [options]
	ponpe [options] <word> <words>...
	ponpe -h | --help
	ponpe -v | --version

Options:
	-h --help                     Show this screen.
	-v --version                  Show version.`
)

func main() {
	os.Exit(Main(os.Args))
}

func Main(argv []string) int {
	parser := &docopt.Parser{}
	args, _ := parser.ParseArgs(doc, argv[1:], Version)
	config := Config{}
	err := args.Bind(&config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	var word []rune
	var marks [][]rune
	for _, orgMark := range config.Words {
		// ダイアクリティカルマークに変換可能かチェック
		mark := []rune(orgMark)
		if err := unicode.ValidateDiaCriticalMark(mark); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
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
	return 0
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
