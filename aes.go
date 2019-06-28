package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

const KEY = "9871267812345mn812345xyz"

//AesEncode aes算法加密
func AesEncode(origin string) string {
	originData := []byte(origin)
	k := []byte(KEY)

	//获取分组密钥
	block, _ := aes.NewCipher(k)
	blockSize := block.BlockSize()

	//数据块进行填充
	originData = Padding(originData, blockSize)

	//加密模式(new Encrypter)
	model := cipher.NewCBCEncrypter(block, k[:blockSize])
	crypted := make([]byte, len(originData))

	//加密
	model.CryptBlocks(crypted, originData)
	return base64.StdEncoding.EncodeToString(crypted)
}

//AesDecode AES算法解密
func AesDecode(crypted string) string {
	cryptByte, _ := base64.StdEncoding.DecodeString(crypted)
	k := []byte(KEY)

	block, _ := aes.NewCipher(k)
	blockSize := block.BlockSize()

	//解密模式(new Decrypter)
	model := cipher.NewCBCDecrypter(block, k[:blockSize])
	origin := make([]byte, len(cryptByte))
	model.CryptBlocks(origin, cryptByte)

	origin = UnPadding(origin)
	return string(origin)
}

//Padding PKCS#7填充数据块
func Padding(data []byte, blockSize int) []byte {
	p := blockSize - len(data)%blockSize
	text := bytes.Repeat([]byte{byte(p)}, p)
	return append(data, text...)
}


//UnPadding PKCS#7去填充
func UnPadding(data []byte) []byte {
	length := len(data)
	unPadding := int(data[length-1])
	return data[:(length-unPadding)]
}
