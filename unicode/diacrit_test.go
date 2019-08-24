package unicode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDiaCriticalMark(t *testing.T) {
	type TestData struct {
		desc    string
		in      []rune
		wantErr bool
	}
	tds := []TestData{
		{desc: "aは変換可能", in: []rune{'a'}},
		{desc: "haは変換可能", in: []rune{'h', 'a'}},
		{desc: "bは変換不可", in: []rune{'b'}, wantErr: true},
	}
	for _, v := range tds {
		t.Run(v.desc, func(t *testing.T) {
			err := ValidateDiaCriticalMark(v.in)
			if v.wantErr {
				assert.Error(t, err, v.desc)
			}
		})
	}
}

func TestToDiaCriticalMark(t *testing.T) {
	type TestData struct {
		desc string
		in   []rune
		want []rune
	}
	tds := []TestData{
		{desc: "aは変換可能", in: []rune{'a'}, want: []rune{'\u0363'}},
		{desc: "haは変換可能", in: []rune{'h', 'a'}, want: []rune{'\u036A', '\u0363'}},
		{desc: "bは変換不可", in: []rune{'b'}, want: []rune{0}},
	}
	for _, v := range tds {
		t.Run(v.desc, func(t *testing.T) {
			got := ToDiaCriticalMark(v.in)
			assert.Equal(t, v.want, got, v.desc)
		})
	}
}
