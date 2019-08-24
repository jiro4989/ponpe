package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	type TestData struct {
		desc string
		in   []string
		want int
	}
	tests := []TestData{
		{desc: "正常系: aaとdd", in: []string{"aa", "dd"}, want: 0},
		{desc: "異常系: aaとbb", in: []string{"aa", "bb"}, want: 2},
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
		inM  []rune
		want string
	}
	tests := []TestData{
		{desc: "wはmより長い", inW: []rune{'a', 'b'}, inM: []rune{'z'}, want: "azb"},
		{desc: "wはmと同じ長さ", inW: []rune{'a', 'b'}, inM: []rune{'z', 'y'}, want: "azby"},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := joinWords(tt.inW, tt.inM)
			assert.Equal(t, tt.want, got, tt.desc)
		})
	}
}
