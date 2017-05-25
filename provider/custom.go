package provider

import (
	"time"
	"math/rand"
	"strconv"
	"fmt"
)

func AirConCMD01(开关机, 模式, 睡眠, 扫风, 风速 int, 温度 string, 灯光, 超强, 上下扫风, 左右扫风, 睡眠4, 睡眠23 int) []byte {
	cmd1 := []byte{1, 45, 45, 50, 25, 151, 64, 1, 8, 1, 0, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}

	cmd1[5] = byte(开关机<<7 + 模式<<4 + 睡眠<<3 + 扫风<<2 + 风速)

	dat6, _ := strconv.Atoi(温度)
	cmd1[6] = byte((dat6 - 16 ) << 4)

	cmd1[7] = byte(1<<3 + 灯光<<1 + 超强)

	cmd1[9] = byte(上下扫风<<4 + 左右扫风)

	cmd1[17] = byte(睡眠4<<7 + 睡眠23<<4)

	return cmd1
}

func acBinGen(cmd1, cmd33, cmd34, cmd35, cmd40, cmd44 []byte, mac string) []byte {
	dat1 := datGen(cmd1)
	dat33 := datGen(cmd33)
	dat34 := datGen(cmd34)
	//dat35 := datGen(cmd35)
	dat40 := datGen(cmd40)
	//dat44 := datGen(cmd44)
	fmt.Println(dat40)
	//dat := make([]byte, len(dat1)+len(dat33)+len(dat34)+len(dat35)+len(dat40))
	dat := make([]byte, len(dat1)+len(dat33)+len(dat34)+len(dat40))
	//dat := make([]byte, len(dat1)+len(dat33)+len(dat34)+len(dat35)+len(dat40)+len(dat44))

	copy(dat[:len(dat1)], dat1)
	copy(dat[len(dat1):len(dat1)+len(dat33)], dat33)
	copy(dat[len(dat1)+len(dat33):len(dat1)+len(dat33)+len(dat34)], dat34)
	//copy(dat[len(dat1)+len(dat33)+len(dat34):len(dat1)+len(dat33)+len(dat34)+len(dat35)], dat35)
	//copy(dat[len(dat1)+len(dat33)+len(dat34)+len(dat35):len(dat1)+len(dat33)+len(dat34)+len(dat35)+len(dat40)], dat40)
	//copy(dat[len(dat1)+len(dat33)+len(dat34)+len(dat35)+len(dat40):], dat44)
	//copy(dat[len(dat1)+len(dat33)+len(dat34)+len(dat35):], dat40)
	copy(dat[len(dat1)+len(dat33)+len(dat34):], dat40)

	fmt.Println(dat)
	bin := binGen(dat, mac)

	copy(bin[22:26], midGen(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10)))

	return bin
}

func midGen(i int) []byte {
	//"10001", "10002", "10003", "11002", "11000"
	mids := make(map[int][]byte)
	mids[0] = []byte{0, 1, 0, 1}
	mids[1] = []byte{0, 1, 0, 2}
	mids[2] = []byte{0, 1, 0, 3}
	mids[3] = []byte{0, 1, 16, 2}
	mids[4] = []byte{0, 1, 16, 0}
	mids[5] = []byte{0, 1, 16, 2}
	mids[6] = []byte{0, 1, 0, 2}
	mids[7] = []byte{0, 1, 0, 1}
	mids[8] = []byte{0, 1, 0, 2}
	mids[9] = []byte{0, 1, 0, 3}
	mids[10] = []byte{0, 1, 16, 0}
	return mids[i]
}
