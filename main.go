package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jiro4989/ponpe/unicode"
)

type (
	ErrorCode int
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
	opts, err := ParseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	os.Exit(int(Main(opts)))
}

func Main(opts *CmdArgs) ErrorCode {
	if opts.Version {
		fmt.Println(Version)
		return errorCodeOk
	}

	if opts.List != "" {
		return cmdList(opts)
	}

	// 引数未指定の時は ponponpain が自動的に設定される
	if len(opts.Args) < 2 {
		opts.Args = []string{"ponponpain", "haraita-i"}
	}

	return cmdJoin(opts)
}

// cmdJoin は2つの入力を結合し、標準出力に出す。
// 引数が'-'という指定のときは、その位置に標準入力から受け取った文字列を埋め込む
// 。'-'指定が存在しないときは標準入力を受け付けない。
func cmdJoin(opts *CmdArgs) ErrorCode {
	if err := setStdinToArgs(opts); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return errorCodeFailedReadingStdin
	}

	var word []rune
	var marks [][]rune
	for _, orgMark := range opts.Args[1:] {
		// 結合文字に変換可能かチェック
		mark := []rune(orgMark)
		if err := unicode.ValidateCombindingCharacterMap(mark); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return errorCodeIllegalAlphabet
		}

		// 結合文字に変換
		mark = unicode.ToCombindingCharacterMap(mark)
		// 結合先の文字よりも、結合文字が多くてはならない
		w, m := deleteOverSize([]rune(opts.Args[0]), mark)

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
func setStdinToArgs(opts *CmdArgs) error {
	if hasStdinArgs(opts.Args) {
		stdinStr, err := readStdin()
		if err != nil {
			return err
		}
		// 受け取った標準入力で上書き
		if opts.Args[0] == "-" {
			opts.Args[0] = stdinStr
		}
		for i := 1; i < len(opts.Args); i++ {
			if opts.Args[i] == "-" {
				opts.Args[i] = stdinStr
			}
		}
	}
	return nil
}

func hasStdinArgs(args []string) bool {
	ws := args
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
			m := mm[i]
			if m == ' ' {
				continue
			}
			s += string(m)
		}
	}
	return s
}

func cmdList(opts *CmdArgs) ErrorCode {
	converter := categories[opts.List]
	for k, v := range converter {
		w, j := string(k), string(v)
		s := fmt.Sprintf("%s  %s %d u%.4x", w, j, v, v)
		fmt.Println(s)
	}
	return errorCodeOk
}
