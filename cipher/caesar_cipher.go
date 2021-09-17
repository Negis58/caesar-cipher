package cipher

import (
	"bytes"
)

type CaesarCipher interface {
	Encrypt(inString string, key int) string
	Decrypt(inString string, key int) string
}

type caesarCipher struct{}

func NewCaesarCipher() CaesarCipher {
	return &caesarCipher{}
}
func (c *caesarCipher) Encrypt(inString string, key int) string {
	var text bytes.Buffer
	for _, char := range inString {
		if char == 32 {
			text.WriteByte(byte(char))
			continue
		}
		cipherChar := getCharStr((getCharPos(char)+key)%26, char)
		text.WriteString(cipherChar)

	}
	return text.String()
}

func (c *caesarCipher) Decrypt(inString string, key int) string {
	var text bytes.Buffer

	for _, char := range inString {
		if char == 32 {
			text.WriteByte(byte(char))
			continue
		}
		mod := (getCharPos(char) - key) % 26
		if mod < 0 {
			mod += 26
		}
		originCipherChar := getCharStr(mod, char)
		text.WriteString(originCipherChar)
	}
	return text.String()
}

func getCharStr(alpPos int, ascii int32) string {
	isLower := ascii >= 97
	if isLower {
		return string(rune('a' - 1 + alpPos))
	}
	return string(rune('A' - 1 + alpPos))
}

func getCharPos(ascii int32) int {
	return int(ascii) - 96
}

