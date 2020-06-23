package koreanutil

import (
	"regexp"
)

// GetJamo 문장을 입력받아 자모 단위로 분리하여 반환
// 한글이 아닌 문자는 그대로 반환
func GetJamo(userInput string) [][]string {
	runeUserInput := []rune(userInput)
	var result [][]string

	ptnKorean := regexp.MustCompile(`[가-힣]`)

	for _, char := range runeUserInput {
		if !ptnKorean.MatchString(string(char)) {
			result = append(result, []string{string(char)})
			continue
		}
		choIdx := (char - BASECODE) / 21 / 28
		jungIdx := ((char - BASECODE) - (choIdx * 21 * 28)) / 28
		jongIdx := (char - BASECODE) - (choIdx * 21 * 28) - (jungIdx * 28)
		result = append(result, []string{string(CHO[(char-BASECODE)/21/28]), string(JUNG[jungIdx]), string(JONG[jongIdx])})
	}
	return result
}
