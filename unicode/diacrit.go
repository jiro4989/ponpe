package unicode

var (
	DiaCriticalMarks = map[rune]rune{
		'a': '\u0363',
		'e': '\u0364',
		'i': '\u0365',
		'o': '\u0366',
		'u': '\u0367',
		'c': '\u0368',
		'd': '\u0369',
		'h': '\u036A',
		'm': '\u036B',
		'r': '\u036C',
		't': '\u036D',
		'v': '\u036E',
		'x': '\u036F',
	}
	// DiaCriticalMarksSupplement はダイアクリティカルマーク補助
	DiaCriticalMarksSupplement = map[rune]rune{
		'q': '\u1DD2',
		'g': '\u1dda',
		'G': '\u1ddb',
		'k': '\u1ddc',
		'L': '\u1dde',
		'R': '\u1de2',
		's': '\u1de4',
		'z': '\u1de6',
		'f': '\u1deb',
		'j': '\u1def', // 微妙だけど
		'l': '\u1ddd',
	}
)
