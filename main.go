package main

import (
	"fmt"

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
	parser := &docopt.Parser{}
	args, _ := parser.ParseArgs(doc, nil, Version)
	config := Config{}
	err := args.Bind(&config)
	if err != nil {
		panic(err)
	}

	// ダイアクリティカルマークへ変換
	var marks []rune
	for _, v := range []rune(config.Words[0]) {
		mark := unicode.DiaCriticalMarks[v]
		marks = append(marks, mark)
	}
	mark := string(marks)

	// 結合先の文字よりも、結合文字が多くてはならない
	word := config.Word
	rWord := []rune(word)
	rMark := []rune(mark)
	if len(rWord) < len(rMark) {
		rMark = rMark[:len(rWord)]
	}

	// 結合して出力
	var s string
	for i := 0; i < len(rWord); i++ {
		w := rWord[i]
		s += string(w)
		if len(rMark) <= i {
			continue
		}
		m := rMark[i]
		s += string(m)
	}
	fmt.Println(s)
}
