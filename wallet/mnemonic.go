package wallet

import (
	"github.com/tyler-smith/go-bip39"
	"github.com/tyler-smith/go-bip39/wordlists"
	"log"
	"regexp"
	"strings"
)

func ConvertMnemonic(mnemonic string) (string, error) {
	var original []string
	var target []string
	if isChinese(mnemonic) {
		original = wordlists.ChineseSimplified
		target = wordlists.English
	} else {
		original = wordlists.English
		target = wordlists.ChineseSimplified
	}
	bip39.SetWordList(original)
	entropy, err := bip39.EntropyFromMnemonic(mnemonic)
	if err != nil {
		return "", err
	}
	log.Printf("%x\n", entropy)
	bip39.SetWordList(target)
	return bip39.NewMnemonic(entropy)
}

func isChinese(mnemonic string) bool {
	words := strings.Fields(mnemonic)
	var exp = regexp.MustCompile("^[\u4e00-\u9fa5]$")
	if exp.MatchString(words[0]) {
		return true
	} else {
		return false
	}
}