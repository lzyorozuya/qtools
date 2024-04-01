package qrandom

import "testing"

func Test1(t *testing.T) {
	t.Log(CharsetCapitalLetters & CharsetCapitalLetters)
	t.Log((CharsetCapitalLetters + CharsetDigit) & CharsetCapitalLetters)
	t.Log((CharsetCapitalLetters + CharsetDigit) & CharsetLowercaseLetter)
}

func TestString(t *testing.T) {
	t.Log(String(CharsetDigit, 128))
	t.Log(String(CharsetLowercaseLetter, 128))
	t.Log(String(CharsetCapitalLetters, 128))
	t.Log(String(CharsetDigit+CharsetLowercaseLetter, 128))
	t.Log(String(CharsetDigit+CharsetCapitalLetters, 128))
	t.Log(String(CharsetLowercaseLetter+CharsetCapitalLetters, 128))
	t.Log(String(CharsetDigit+CharsetLowercaseLetter+CharsetCapitalLetters, 128))

	t.Log(String(CharsetSpecial1, 128))
	t.Log(String(CharsetSpecial1+CharsetSpecial2, 128))
	t.Log(String(CharsetSpecial1+CharsetSpecial2+CharsetSpecial3, 128))
	t.Log(String(CharsetSpecial1+CharsetSpecial2+CharsetSpecial3+CharsetSpecial4, 128))
	t.Log(String(CharsetDigit+CharsetBlank, 128))
}
