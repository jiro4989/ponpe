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
		{desc: "正常系: join: aaとdd", in: []string{"./bin/ponpe", "join", "aa", "dd"}, want: errorCodeOk},
		{desc: "正常系: join: ponponpainとharaita-i", in: []string{"ponpe", "join", "ponponpain", "haraita-i"}, want: errorCodeOk},
		{desc: "正常系: join: aaとddとcc", in: []string{"ponpe", "join", "aa", "cc", "dd"}, want: errorCodeOk},
		{desc: "異常系: join: aaとzz", in: []string{"ponpe", "join", "aa", "zz"}, want: errorCodeIllegalAlphabet},
		{desc: "正常系: join: abcdeとabcde(キリル文字)", in: []string{"ponpe", "join", "abcde", "abcde"}, want: errorCodeOk},
		{desc: "正常系: j: aaとdd (joinのエイリアス)", in: []string{"ponpe", "j", "aa", "dd"}, want: errorCodeOk},
		{desc: "正常系: list: all", in: []string{"ponpe", "list", "all"}, want: errorCodeOk},
		{desc: "正常系: list: a", in: []string{"ponpe", "list", "a"}, want: errorCodeOk},
		{desc: "正常系: list: diacritical_mark", in: []string{"ponpe", "list", "diacritical_mark"}, want: errorCodeOk},
		{desc: "正常系: list: dm", in: []string{"ponpe", "list", "dm"}, want: errorCodeOk},
		{desc: "正常系: list: cyrillic_alphabets", in: []string{"ponpe", "list", "cyrillic_alphabets"}, want: errorCodeOk},
		{desc: "正常系: list: ca", in: []string{"ponpe", "list", "ca"}, want: errorCodeOk},
		{desc: "正常系: l: ca (listのエイリアス)", in: []string{"ponpe", "l", "ca"}, want: errorCodeOk},
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
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := joinWords(tt.inW, tt.inM...)
			assert.Equal(t, tt.want, got, tt.desc)
		})
	}
}
