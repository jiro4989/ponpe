package unicode

// キリル文字
// https://ja.wikipedia.org/wiki/%E3%82%AD%E3%83%AA%E3%83%AB%E6%96%87%E5%AD%97#Unicode_%E5%8F%8E%E9%8C%B2%E4%BD%8D%E7%BD%AE
// キリル文字一覧: https://ja.wikipedia.org/wiki/%E3%82%AD%E3%83%AA%E3%83%AB%E6%96%87%E5%AD%97%E4%B8%80%E8%A6%A7
// Unicode一覧表: http://www.shurey.com/js/works/unicode.html
var (
	// CyrillicAlphabets はキリル文字。
	CyrillicAlphabets = map[rune]rune{
		'K': '\u2DE6',
		'M': '\u2DE8',
		'H': '\u2DE9',
		'O': '\u2DEA',
		'n': '\u2DEB',
		'p': '\u2DEC',
		'C': '\u2DED',
		'T': '\u2DEE',
		'X': '\u2DEF',
		'U': '\u2DF0',
		'Y': '\u2DF1',
		'w': '\u2DF2',
		'E': '\uA674',
		'N': '\uA675',
		'y': '\uA677',
		'b': '\uA67A',
		'W': '\uA67B',
	}
)
