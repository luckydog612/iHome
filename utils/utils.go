package utils

import (
	"crypto/sha256"
	"encoding/json"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/ripemd160"
	"iHome/models"
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

// 将字符串转化为rears
func StringToAreas(data string) []models.Area {
	areas := make([]models.Area, 0)
	err := json.Unmarshal([]byte(data), &areas)
	if err != nil {
		beego.Error("unmarshal areas err: ", err)
	}
	return areas
}
