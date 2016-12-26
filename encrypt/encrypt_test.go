package encrypt

import (
	"fmt"
	"log"
	"github.com/wanderself/sliceTest/encrypt"
)

func test() {
	//aircon
	//org := []byte{126, 126, 47, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 126, 126, 47, 51, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 126, 126, 47, 52, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 126, 126, 47, 53, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	//ricecooker
	org := []byte{126, 126, 3, 149, 170, 5, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println("元数据：  ", org)

	cipherText := AesEncrypt(org, encrypt.GetKey(1))
	fmt.Println("加密数据： ", cipherText)

	restoreText, err := DecryptCommonKey(cipherText)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("还原数据： ", restoreText)
}
