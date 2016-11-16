package encrypt

import (
	"errors"
	"fmt"
	"crypto/aes"
	"bytes"
)

const (
	//`互联网通讯统一密钥`
	UNIFORM_AES_KEY_COMM = "a3K8Bx%2r8Y7#xDh"

	//服务器通讯统一密钥  服务器之间的加解密公共密钥
	UNIFORM_AES_KEY_SVR = "54*OpC3$x9y-1Dak"

	//用于加密数据库用户登陆密码
	UNIFORM_AES_KEY_DBUSER = "99d#sf%ae85*OP1D"
)
const (
	EncryptKey_Device  = iota //0 设备秘钥加密
	EncryptKey_Uniform        //1 公共秘钥加密
	EncryptKey_Server         //2 服务器秘钥加密

)

func DecryptCommonKey(dat []byte) ([]byte, error) {

	if len(dat)%16 != 0 {
		return nil, errors.New(fmt.Sprintf("decryptCommonKey 不是合法aes加密包 len=%d", len(dat)))
	}

	buf, err := AesDecryptNonBase64(dat, EncryptKey_Uniform)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func AesEncrypt(data []byte, key string) []byte {

	aesCipher, _ := aes.NewCipher([]byte(key))
	encrypter := NewECBEncrypter(aesCipher)

	blockSize := aesCipher.BlockSize()
	origData := PKCS5Padding(data, blockSize)

	dest := make([]byte, len(origData))

	encrypter.CryptBlocks(dest, origData)

	return dest
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func GetKey(encryptKeyType int) string {
	var key string
	switch encryptKeyType {
	case 0:
		key = "unkown"
	case 1:
		key = UNIFORM_AES_KEY_COMM
	case 2:
		key = UNIFORM_AES_KEY_SVR
	}
	return key
}

func AesDecryptNonBase64(data []byte, encryptKeyType int) ([]byte, error) {

	aesCipher, _ := aes.NewCipher([]byte(GetKey(encryptKeyType)))
	decrypter := NewECBDecrypter(aesCipher)

	if len(data)%decrypter.BlockSize() != 0 {
		return nil, errors.New("数据长度错误.")
	}
	//blockSize := aesCipher.BlockSize()

	dest := make([]byte, len(data))

	decrypter.CryptBlocks(dest, data)

	dest = PKCS5UnPadding(dest)

	return dest, nil

}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}