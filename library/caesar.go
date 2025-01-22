package library

import "strings"

var hashSalt []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func CaesarEncrypted(codeText string, status, shift int) (string, error) {
	var baseCode strings.Builder

	for _, text := range codeText {
		index := -1

		for i, val := range hashSalt {
			if val == text {
				index = i
				break
			}
		}

		if index != -1 {
			if status == 1 {
				baseCode.WriteRune(hashSalt[(index+shift)%len(hashSalt)])
			} else {
				baseCode.WriteRune(hashSalt[(index-shift+len(hashSalt))%len(hashSalt)])
			}
		} else {
			baseCode.WriteRune(text)
		}
	}
	return baseCode.String(), nil
}
