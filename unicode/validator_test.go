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
	tests := []TestData{
		{desc: "aは変換可能", in: []rune{'a'}},
		{desc: "haは変換可能", in: []rune{'h', 'a'}},
		{desc: "bは変換不可", in: []rune{'b'}, wantErr: true},
		{desc: "半角スペースはスキップ", in: []rune{' '}},
		{desc: "全角スペースはエラー", in: []rune{'　'}, wantErr: true},
		{desc: "タブ文字はエラー", in: []rune{'\t'}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			err := ValidateDiaCriticalMark(tt.in)
			if tt.wantErr {
				assert.Error(t, err, tt.desc)
			}
		})
	}
}

func TestValidateCyrillicAlphabets(t *testing.T) {
	type TestData struct {
		desc    string
		in      []rune
		wantErr bool
	}
	tests := []TestData{
		{desc: "rは変換可能", in: []rune{'r'}},
		{desc: "aは変換不可", in: []rune{'a'}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			err := ValidateCyrillicAlphabets(tt.in)
			if tt.wantErr {
				assert.Error(t, err, tt.desc)
			}
		})
	}
}

func TestValidateCombindingCharacterMap(t *testing.T) {
	type TestData struct {
		desc    string
		in      []rune
		wantErr bool
	}
	tests := []TestData{
		{desc: "aは変換可能", in: []rune{'a'}},
		{desc: "rは変換可能", in: []rune{'r'}},
		{desc: "あは変換不可", in: []rune{'あ'}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			err := ValidateCombindingCharacterMap(tt.in)
			if tt.wantErr {
				assert.Error(t, err, tt.desc)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	type TestData struct {
		desc    string
		inW     []rune
		inM     map[rune]rune
		wantErr bool
	}
	tests := []TestData{
		{desc: "マップが持つキーは変換可能", inW: []rune{'a'}, inM: map[rune]rune{'a': 1}},
		{desc: "マップが持たないキーは変換不可", inW: []rune{'b'}, inM: map[rune]rune{'a': 1}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			err := validate(tt.inW, tt.inM, "")
			if tt.wantErr {
				assert.Error(t, err, tt.desc)
			}
		})
	}
}
