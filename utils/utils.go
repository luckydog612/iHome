package utils

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"log"
	"strconv"
	"strings"
)

func HashName(avatarByte []byte) []byte {
	publicSHA256 := sha256.Sum256(avatarByte)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)
	return publicRIPEMD160
}

func UniCodeToString(str string) string {
	sUnicodev := strings.Split(str, "\\u")
	var context string
	for _, v := range sUnicodev {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			panic(err)
		}
		context += fmt.Sprintf("%c", temp)
	}
	return context
}
