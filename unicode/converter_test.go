package unicode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDiaCriticalMark(t *testing.T) {
	type TestData struct {
		desc string
		in   []rune
		want []rune
	}
	tests := []TestData{
		{desc: "aは変換可能", in: []rune{'a'}, want: []rune{'\u0363'}},
		{desc: "haは変換可能", in: []rune{'h', 'a'}, want: []rune{'\u036A', '\u0363'}},
		{desc: "bは変換不可", in: []rune{'b'}, want: []rune{0}},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := ToDiaCriticalMark(tt.in)
			assert.Equal(t, tt.want, got, tt.desc)
		})
	}
}
