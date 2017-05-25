package provider

import "wander/test/DataProvider-master/encrypt"

func Purifier(date string, cmd byte) []byte {
	evt := make([]byte, 0)
	binHead := []byte{102, 103, 0, 17, 32, 89, 160, 181, 0, 121, 0, 0, 0, 0, 0, 0, 244, 145, 30, 16, 175, 239, 0, 0, 134, 0, 0, 0, 0, 1, 4, 0, 16, 9, 12, 9, 24, 22, 2}
	copy(binHead[16:22], macGen(mac))
	copy(binHead[32:38], dateCus(date))
	copy(binHead[30:31], append(evt, cmd))
	dat44 := datGen([]byte{68, 0, 1})
	dat92 := datGen([]byte{146, 0, 25, 3, 0, 0})
	dat93 := datGen([]byte{147, 1, 0})
	dat94 := datGen([]byte{148, 1, 2, 200, 0, 5, 6, 7, 8, 9, 188, 2, 2, 39, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3})
	dat95 := datGen([]byte{149, 5, 0})
	dat := make([]byte, len(dat44)+len(dat92)+len(dat93)+len(dat94)+len(dat95))
	copy(dat[:len(dat44)], dat44)
	copy(dat[len(dat44):len(dat44)+len(dat92)], dat92)
	copy(dat[len(dat44)+len(dat92):len(dat44)+len(dat92)+len(dat93)], dat93)
	copy(dat[len(dat44)+len(dat92)+len(dat93):len(dat44)+len(dat92)+len(dat93)+len(dat94)], dat94)
	copy(dat[len(dat44)+len(dat92)+len(dat93)+len(dat94):], dat95)

	datLen := getLen(len(dat))
	binHead[3] = byte(datLen)

	//bin := binGen(dat)
	bundle := encrypt.AesEncrypt(dat, encrypt.GetKey(1))
	bin := make([]byte, len(binHead)+len(bundle))

	copy(bin, binHead)
	copy(bin[len(binHead):], bundle)
	wpMid := []byte{0, 130, 131, 1}
	copy(bin[22:26], wpMid)

	return bin
}
