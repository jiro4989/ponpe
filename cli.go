package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jiro4989/ponpe/unicode"
)

type CmdArgs struct {
	List string
	Args []string
}

var (
	categories = map[string]map[rune]rune{
		"diacritical_mark":   unicode.DiaCriticalMarks,
		"dm":                 unicode.DiaCriticalMarks,
		"cyrillic_alphabets": unicode.CyrillicAlphabets,
		"ca":                 unicode.CyrillicAlphabets,
		"all":                unicode.CombindingCharacterMap,
	}
)

func ParseArgs() (*CmdArgs, error) {
	opts := CmdArgs{}

	flag.Usage = flagHelpMessage
	flag.StringVar(&opts.List, "l", "", "print available alphabets. [all|diacritical_mark|dm|cyrillic_alphabets|ca]")
	flag.Parse()
	opts.Args = flag.Args()

	if err := opts.Validate(); err != nil {
		return nil, err
	}

	return &opts, nil
}

func flagHelpMessage() {
	cmd := os.Args[0]
	fmt.Fprintln(os.Stderr, fmt.Sprintf("%s はpͪoͣnⷢpͣoꙶnͭpͣa͡iꙶnを再現するためのアルファベット結合ユーティリティです。", cmd))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s <word> <words>...", cmd))
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s -l <category>", cmd))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Examples:")
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s ponponpain haraita-i", cmd))
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s abcd dddd aaaa tttt eeee", cmd))
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  echo abcd | %s - date", cmd))
	fmt.Fprintln(os.Stderr, fmt.Sprintf("  %s -l all", cmd))
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Available words:")
	fmt.Fprintln(os.Stderr, fmt.Sprintf(`  全ての文字が使用可能なわけではありません。文字列の結合が可能かどうかは、入力
  した文字列に紐づくUnicode結合文字が存在するかどうかで決まります。

  このコマンドはアルファベットと、そのアルファベット類似した形のUnicode結合文
  字を紐づけることで、アルファベットをUnicode結合文字に変換して結合しています。

  よって、Unicode結合文字側に存在しない（＝マッピングされていない）文字を指定
  しても結合できません。

  最低限アルファベット小文字はマッピングしてありますが、アルファベット大文字は
  部分的にしかマッピングしていません。使用可能な文字の一覧については、以下のコ
  マンドで確認してください。

    ponpe -l all

  また、使用している文字が特殊な文字であるため、フォントによっては表示されない
  場合があることをご理解ください。`))
	fmt.Fprintln(os.Stderr, "Options:")

	flag.PrintDefaults()
}

func (c *CmdArgs) Validate() error {
	if c.List != "" {
		if _, ok := categories[c.List]; !ok {
			return fmt.Errorf("'-l' parameter was illegal string. see '-h'.")
		}
		return nil
	}

	return nil
}
