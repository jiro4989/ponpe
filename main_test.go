package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	type TestData struct {
		desc string
		in   []string
		want ErrorCode
	}
	tests := []TestData{
		{desc: "正常系: aaとdd", in: []string{"./bin/ponpe", "aa", "dd"}, want: errorCodeOk},
		{desc: "正常系: ponponpainとharaita-i", in: []string{"ponpe", "ponponpain", "haraita-i"}, want: errorCodeOk},
		{desc: "正常系: aaとddとcc", in: []string{"ponpe", "aa", "cc", "dd"}, want: errorCodeOk},
		{desc: "正常系: a..zとa..z", in: []string{"ponpe", "abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz"}, want: errorCodeOk},
		{desc: "異常系: aaとあ", in: []string{"ponpe", "aa", "あ"}, want: errorCodeIllegalAlphabet},
		{desc: "正常系: abcdeとabcde(キリル文字)", in: []string{"ponpe", "abcde", "abcde"}, want: errorCodeOk},
		{desc: "正常系: 空白文字はスキップ", in: []string{"ponpe", "abc", "a a"}, want: errorCodeOk},
		{desc: "正常系: --list: all", in: []string{"ponpe", "--list", "all"}, want: errorCodeOk},
		{desc: "正常系: --list: diacritical_mark", in: []string{"ponpe", "--list", "diacritical_mark"}, want: errorCodeOk},
		{desc: "正常系: --list: dm", in: []string{"ponpe", "--list", "dm"}, want: errorCodeOk},
		{desc: "正常系: --list: cyrillic_alphabets", in: []string{"ponpe", "--list", "cyrillic_alphabets"}, want: errorCodeOk},
		{desc: "正常系: --list: ca", in: []string{"ponpe", "--list", "ca"}, want: errorCodeOk},
		{desc: "正常系: -l: ca (listのエイリアス)", in: []string{"ponpe", "-l", "ca"}, want: errorCodeOk},
		{desc: "正常系: 引数なしの場合は ponponpain haraita-i が自動挿入される", in: []string{"ponpe"}, want: errorCodeOk},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := Main(tt.in)
			assert.Equal(t, tt.want, got, tt.desc)
		})
	}
}

func TestDeleteOverSize(t *testing.T) {
	type TestData struct {
		desc  string
		inW   []rune
		inM   []rune
		wantW []rune
		wantM []rune
	}
	tests := []TestData{
		{desc: "wはmより長い", inW: []rune{'a', 'b'}, inM: []rune{'z'}, wantW: []rune{'a', 'b'}, wantM: []rune{'z'}},
		{desc: "wはmより短い", inW: []rune{'a'}, inM: []rune{'z', 'y'}, wantW: []rune{'a'}, wantM: []rune{'z'}},
		{desc: "wはmと同じ長さ", inW: []rune{'a', 'b'}, inM: []rune{'z', 'y'}, wantW: []rune{'a', 'b'}, wantM: []rune{'z', 'y'}},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotW, gotM := deleteOverSize(tt.inW, tt.inM)
			assert.Equal(t, tt.wantW, gotW, tt.desc)
			assert.Equal(t, tt.wantM, gotM, tt.desc)
		})
	}
}

func TestJoinWords(t *testing.T) {
	type TestData struct {
		desc string
		inW  []rune
		inM  [][]rune
		want string
	}
	tests := []TestData{
		{desc: "wはmより長い", inW: []rune{'a', 'b'}, inM: [][]rune{{'z'}}, want: "azb"},
		{desc: "wはmと同じ長さ", inW: []rune{'a', 'b'}, inM: [][]rune{{'z', 'y'}}, want: "azby"},
		{desc: "mが3個", inW: []rune{'a', 'b'}, inM: [][]rune{{'z', 'y'}, {'A', 'B'}, {'C', 'D'}}, want: "azACbyBD"},
		{desc: "mが0個", inW: []rune{'a', 'b'}, inM: [][]rune{}, want: "ab"},
		{desc: "半角スペースはスキップ", inW: []rune{'a', 'b'}, inM: [][]rune{{' ', 'y'}}, want: "aby"},
		{desc: "半角スペースはスキップ", inW: []rune{'a', 'b', 'c'}, inM: [][]rune{{' ', ' ', 'y'}}, want: "abcy"},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := joinWords(tt.inW, tt.inM...)
			assert.Equal(t, tt.want, got, tt.desc)
		})
	}
}
