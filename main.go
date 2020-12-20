package main

import (
	"flag"
	"github.com/AllenLiuxt/gotools/wallet"
	"log"
)

func main() {
	mnemonic := flag.String("mnemonic", "", "Input your mnemonics")
	flag.Parse()
	if *mnemonic != "" {
		convert(*mnemonic)
	} else {
		log.Println("Missing mnemonics")
	}
}

func convert(mnemonic string) {
	result, err := wallet.ConvertMnemonic(mnemonic)
	if err != nil {
		log.Println("ConvertMnemonic error:", err.Error())
		return
	}
	log.Println("Converted Mnemonic:", result)

	return
}


