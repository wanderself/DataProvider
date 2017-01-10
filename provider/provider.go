package provider

import (
	"github.com/wanderself/DataProvider/encrypt"
	"fmt"
	"strconv"
	"time"
)

var (
	mac string = "112233445566"

	RC = []byte{149, 124, 5}
	STV_cmd1 = []byte{1, 1, 45, 55, 0, 1}
	STV_cmd31 = []byte{49, 1, 45, 55, 0, 1, 0}
	STV_cmd34 = []byte{52, 10, 14, 128, 5, 127, 7, 0, 0, 0, 0, 0, 0, 105, 1, 90, 0, 150, 0, 105, 1, 90, 0, 150, 0, 144, 1, 30, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 6, 0}
	AC_cmd1 = []byte{1, 45, 45, 50, 25, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}
	AC_cmd33 = []byte{51, 46, 46, 53, 25, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 33, 0, 1}
	AC_cmd34 = []byte{52, 47, 47, 54, 25, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 0, 0, 1, 0, 34, 0, 1}
	//AC_cmd35 = []byte{53, 48, 48, 55, 25, 1, 0, 1, 0, 1, 0, 1, 255, 1, 0, 1, 0, 255, 0, 1, 32, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 255, 0, 1, 0, 1, 0, 255, 0, 1, 0, 1, 0, 35, 0, 1}
	AC_cmd35 = []byte{53, 48, 48, 55, 25, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 255, 1, 0, 1, 0, 35, 0, 1}
	AC_cmd44 = []byte{68, 2, 16, 1, 0, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 32, 33, 1, 2, 3, 0}
	//AC_cmd44 = []byte{68, 2, 16, 1, 0, 49, 50, 51, 52, 53, 54, 55, 56, 57, 48, 65, 66, 67, 95, 95, 1, 2, 3, 0}
	test = []byte{159, 126, 154, 222, 21, 32, 57, 177, 230, 147, 190, 11, 167, 190, 218, 186, 32, 124, 187, 5, 204, 119, 58, 203, 123, 164, 182, 22, 223, 102, 103, 0}
)

func Hatcher(code int) []byte {

	//device mid
	stvMid := []byte{0, 0, 134, 0}
	acMid := []byte{0, 1, 16, 2}
	rcMid := []byte{0, 130, 132, 1}

	switch code {

	case 3:
		dat := datGen(RC)
		bin := binGen(dat)
		copy(bin[22:26], rcMid)
		return bin
	case 1:
		dat1 := datGen(STV_cmd1)
		dat31 := datGen(STV_cmd31)
		dat34 := datGen(STV_cmd34)
		dat := make([]byte, len(dat1) + len(dat31) + len(dat34))
		copy(dat[:len(dat1)], dat1)
		copy(dat[len(dat1):], dat31)
		copy(dat[len(dat1) + len(dat31):], dat34)
		bin := binGen(dat)
		copy(bin[22:26], stvMid)
		fmt.Println(bin)
		return bin
	case 0:
		dat1 := datGen(AC_cmd1)
		dat33 := datGen(AC_cmd33)
		dat34 := datGen(AC_cmd34)
		dat35 := datGen(AC_cmd35)
		dat44 := datGen(AC_cmd44)
		dat := make([]byte, len(dat1) + len(dat33) + len(dat34) + len(dat35) + len(dat44))

		copy(dat[:len(dat1)], dat1)
		copy(dat[len(dat1):len(dat1) + len(dat33)], dat33)
		copy(dat[len(dat1) + len(dat33) :len(dat1) + len(dat33) + len(dat34)], dat34)
		copy(dat[len(dat1) + len(dat33) + len(dat34):len(dat1) + len(dat33) + len(dat34) + len(dat35)], dat35)
		copy(dat[len(dat1) + len(dat33) + len(dat34) + len(dat35):], dat44)
		bin := binGen(dat)

		fmt.Println(bin)
		copy(bin[22:26], acMid)
		//copy(bin[16:22], mac)
		return bin
	default:
		jsn := []byte("{\"t\":\"notify\",\"mac\":\"2059a0b4214b\",\"mid\":\"30000\",\"tm\":\"20161230125211\",\"evt\":3,\"code\":21,\"p\":\"Pow\",\"o\":\"1\",\"v\":\"0\",\"msg\":\"sd\"}\n")
		return jsn
	}

}

func binGen(dat []byte) []byte {
	//binHead := []byte{102, 103, 0, 17, 32, 89, 160, 181, 0, 121, 0, 0, 0, 0, 0, 0, 32, 89, 160, 180, 0, 19, 0, 0, 134, 0, 0, 0, 0, 1, 0, 0, 16, 9, 12, 9, 24, 22, 2}
	binHead := []byte{102, 103, 0, 17, 32, 89, 160, 181, 0, 121, 0, 0, 0, 0, 0, 0, 244, 145, 30, 16, 175, 239, 0, 0, 134, 0, 0, 0, 0, 1, 4, 0, 16, 9, 12, 9, 24, 22, 2}

	copy(binHead[32:38], tmGen())
	copy(binHead[16:22], macGen(mac))
	datLen := getLen(len(dat))
	binHead[3] = byte(datLen)
	bundle := encrypt.AesEncrypt(dat, encrypt.GetKey(1))
	bin := make([]byte, len(binHead) + len(bundle))

	copy(bin, binHead)
	copy(bin[len(binHead):], bundle)

	return bin
}

func datGen(datBody []byte) []byte {
	datHead := []byte{126, 126}
	datLen := len(datBody) + 3
	//fmt.Println(len(datBody))
	//dat := make([]byte, getLen(datLen))
	dat := make([]byte, datLen)
	//fmt.Println(dat)
	copy(dat[0:2], datHead)
	dat[2] = byte(len(datBody) + 1)
	//copy(dat[3:datLen], datBody)
	copy(dat[3:datLen], datBody)
	return dat
}

func getLen(len int) int {
	if len <= 16 {
		return 16
	} else if len % 16 == 0 {
		return len
	} else {
		len = (len / 16 + 1) * 16
		return len
	}
}

func tmGen() []byte {

	tm := make([]byte, 6)

	Y, _ := strconv.Atoi(time.Now().String()[2:4])

	M, _ := strconv.Atoi(time.Now().String()[5:7])

	D, _ := strconv.Atoi(time.Now().String()[8:10])

	h, _ := strconv.Atoi(time.Now().String()[11:13])

	m, _ := strconv.Atoi(time.Now().String()[14:16])

	s, _ := strconv.Atoi(time.Now().String()[17:19])

	tm[0] = byte(Y)

	tm[1] = byte(M)

	tm[2] = byte(D)

	tm[3] = byte(h)

	tm[4] = byte(m)

	tm[5] = byte(s)

	return tm

}

func macGen(s string) []byte {

	mac := make([]byte, 6)

	m1, _ := strconv.ParseInt(s[:2], 16, 16)

	m2, _ := strconv.ParseInt(s[2:4], 16, 16)

	m3, _ := strconv.ParseInt(s[4:6], 16, 16)

	m4, _ := strconv.ParseInt(s[6:8], 16, 16)

	m5, _ := strconv.ParseInt(s[8:10], 16, 16)

	m6, _ := strconv.ParseInt(s[10:12], 16, 16)

	mac[0] = byte(m1)

	mac[1] = byte(m2)

	mac[2] = byte(m3)

	mac[3] = byte(m4)

	mac[4] = byte(m5)

	mac[5] = byte(m6)

	return mac

}




//dat := []byte{102, 103, 0, 32, 32, 89, 160, 181, 0, 121, 0, 0, 0, 0, 0, 0, 32, 89, 160, 181, 0, 121, 0, 0, 134, 0, 0, 0, 0, 1, 0, 0, 16, 9, 12, 9, 24, 22, 2, 194, 186, 186, 167, 155, 104, 142, 138, 230, 2, 238, 212, 3, 18, 51, 109, 155, 65, 217, 191, 11, 57, 210, 238, 92, 57, 72, 225, 114, 245, 190, 159}

//hf content 198, 177, 248, 14, 120, 110, 228, 181, 251, 38, 142, 146, 169, 4, 171, 65, 218, 2, 87, 110, 118, 62, 14, 28, 221, 49, 234, 73, 199, 41, 224, 8, 40, 233, 12, 8, 110, 15, 207, 51, 236, 65, 23, 191, 112, 41, 33, 45, 140, 182, 130, 249, 89, 15, 82, 218, 142, 43, 62, 73, 146, 85, 77, 103, 181, 115, 123, 37, 71, 253, 165, 165, 249, 21, 35, 178, 54, 44, 158, 183

// head 102, 103, 0, 80, 32, 89, 160, 181, 0, 121, 0, 0, 0, 0, 0, 0, 32, 89, 160, 181, 0, 121, 0, 0, 134, 0, 0, 0, 0, 1, 0, 0, 16, 9, 12, 9, 24, 22, 2,
//ac content 6, 159, 130, 165, 170, 74, 232, 38, 210, 84, 29, 246, 106, 52, 27, 121, 202, 138, 36, 193, 158, 0, 45, 198, 180, 150, 135, 105, 77, 202, 154, 85, 202, 138, 36, 193, 158, 0, 45, 198, 180, 150, 135, 105, 77, 202, 154, 85, 132, 247, 246, 172, 31, 122, 19, 177, 30, 124, 147, 33, 169, 137, 138, 87, 195, 114, 14, 109, 97, 193, 169, 39, 224, 3, 146, 93, 178, 196, 128, 116, 195, 114, 14, 109, 97, 193, 169, 39, 224, 3, 146, 93, 178, 196, 128, 116, 117, 242, 216, 229, 232, 202, 162, 32, 146, 247, 116, 20, 180, 39, 7, 152, 202, 138, 36, 193, 158, 0, 45, 198, 180, 150, 135, 105, 77, 202, 154, 85, 202, 138, 36, 193, 158, 0, 45, 198, 180, 150, 135, 105, 77, 202, 154, 85, 80, 204, 47, 249, 226, 173, 248, 232, 185, 115, 54, 152, 135, 221, 147, 130, 195, 114, 14, 109, 97, 193, 169, 39, 224, 3, 146, 93, 178, 196, 128, 116, 195, 114, 14, 109, 97, 193, 169, 39, 224, 3, 146, 93, 178, 196, 128, 116, 104, 109, 67, 179, 35, 236, 220, 202, 238, 44, 151, 99, 94, 4, 30, 216

// head 102, 103, 0, 17, 32, 89, 160, 181, 0, 121, 0, 0, 0, 0, 0, 0, 32, 89, 160, 181, 0, 121, 0, 0, 134, 0, 0, 0, 0, 1, 0, 0, 16, 9, 12, 9, 24, 22, 2,
//rc content 194, 186, 186, 167, 155, 104, 142, 138, 230, 2, 238, 212, 3, 18, 51, 109, 155, 65, 217, 191, 11, 57, 210, 238, 92, 57, 72, 225, 114, 245, 190, 159
