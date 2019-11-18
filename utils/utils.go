package utils

import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"log"
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
