package qrandom

import (
	"math/rand"
	"strings"
	"sync"
)

type CharsetType uint16

const (
	CharsetDigit           CharsetType = 1 << iota // 数字
	CharsetLowercaseLetter CharsetType = 1 << iota // 小写字母
	CharsetCapitalLetters  CharsetType = 1 << iota // 大写字母
	CharsetSpecial1        CharsetType = 1 << iota // 特殊符号1 !"#$%&'()*+,-./
	CharsetSpecial2        CharsetType = 1 << iota // 特殊符号2 :;<=>?@
	CharsetSpecial3        CharsetType = 1 << iota // 特殊符号3 [\]^_`
	CharsetSpecial4        CharsetType = 1 << iota // 特殊符号4 {|}~
	CharsetBlank           CharsetType = 1 << iota // 空格
)

var (
	charsetType2charset = map[CharsetType]string{
		CharsetDigit:           "0123456789",
		CharsetLowercaseLetter: "qwertyuiopasdfghjklzxcvbnm",
		CharsetCapitalLetters:  "QWERTYUIOPASDFGHJKLZXCVBNM",
		CharsetSpecial1:        "!\"#$%&'()*+,-./",
		CharsetSpecial2:        ":;<=>?@",
		CharsetSpecial3:        "[\\]^_`",
		CharsetSpecial4:        "{|}~",
		CharsetBlank:           " ",
	}
	r sync.RWMutex
)

func GetCharset(t CharsetType) string {
	if s := getCharset(t); s != "" {
		return s
	}

	sb := strings.Builder{}
	for charsetType, s := range charsetType2charset {
		if t&charsetType == charsetType {
			sb.WriteString(s)
		}
	}
	charset := sb.String()
	setCharset(t, charset)
	return charset
}

func getCharset(t CharsetType) string {
	r.RLock()
	defer r.RUnlock()
	return charsetType2charset[t]
}

func setCharset(t CharsetType, charset string) {
	r.Lock()
	defer r.Unlock()
	charsetType2charset[t] = charset
	return
}

func String(t CharsetType, length int) string {
	return StringFromCharset(GetCharset(t), length)
}

func StringFromCharset(charset string, length int) string {
	if length <= 0 {
		return ""
	}
	result := make([]byte, length)
	l := len(charset)
	for i := range result {
		result[i] = charset[rand.Intn(l)]
	}
	return string(result)
}
