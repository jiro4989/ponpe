package main

import (
	"bufio"
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
	doc = `ponpe はpͪoͣnⷢpͣoꙶnͭpͣa͡iꙶnを再現するためのアルファベット結合ユーティリティです。

Usage:
	ponpe [options] <word> <words>...
	ponpe [-l | --list] (all | diacritical_mark | dm | cyrillic_alphabets | ca)
	ponpe -h | --help
	ponpe -v | --version

Examples:
	ponpe ponponpain haraita-i
	ponpe abcd dddd aaaa tttt eeee
	echo abcd | ponpe - date
	ponpe --list all

Available words:
	全ての文字が使用可能なわけではありません。文字列の結合が可能かどうかは、入力
	した文字列に紐づくUnicode結合文字が存在するかどうかで決まります。

	このコマンドはアルファベットと、そのアルファベット類似した形のUnicode結合文
	字を紐づけることで、アルファベットをUnicode結合文字に変換して結合しています。

	よって、Unicode結合文字側に存在しない（＝マッピングされていない）文字を指定
	しても結合できません。

	最低限アルファベット小文字はマッピングしてありますが、アルファベット大文字は
	部分的にしかマッピングしていません。使用可能な文字の一覧については、以下のコ
	マンドで確認してください。

		ponpe --list all

	また、使用している文字が特殊な文字であるため、フォントによっては表示されない
	場合があることをご理解ください。

Options:
	-h --help       このヘルプを出力する。
	-v --version    バージョン情報を出力する。`
)

const (
	errorCodeOk ErrorCode = iota
	errorCodeFailedBinding
	errorCodeFailedReadingStdin
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

// cmdJoin は2つの入力を結合し、標準出力に出す。
// 引数が'-'という指定のときは、その位置に標準入力から受け取った文字列を埋め込む
// 。'-'指定が存在しないときは標準入力を受け付けない。
func cmdJoin(config Config) ErrorCode {
	if err := setStdinToArgs(&config); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return errorCodeFailedReadingStdin
	}

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

// setStdinToArgs は標準入力を受け取る指定があるときだけハイフンの引数の位置に標
// 準入力を埋め込む。
func setStdinToArgs(config *Config) error {
	if hasStdinArgs(*config) {
		stdinStr, err := readStdin()
		if err != nil {
			return err
		}
		// 受け取った標準入力で上書き
		if config.Word == "-" {
			config.Word = stdinStr
		}
		for i := 0; i < len(config.Words); i++ {
			if config.Words[i] == "-" {
				config.Words[i] = stdinStr
			}
		}
	}
	return nil
}

func hasStdinArgs(config Config) bool {
	word, words := config.Word, config.Words
	ws := []string{word}
	ws = append(ws, words...)
	for _, w := range ws {
		if w == "-" {
			return true
		}
	}
	return false
}

func readStdin() (string, error) {
	sc := bufio.NewScanner(os.Stdin)
	var s string
	for sc.Scan() {
		s = sc.Text()
		break
	}
	if err := sc.Err(); err != nil {
		return "", err
	}
	return s, nil
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
