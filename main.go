package main

import (
	"wander/test/DataProvider-master/connector"
	"strconv"
	"fmt"
	"wander/test/DataProvider-master/provider"
	"time"
	"math/rand"
	"log"
)

const (
	AIRCON = 0 //air conditioner
	STOVE  = 1 //hang stove
	FRIDGE = 2 //fridge
	RCKR   = 3 //rice cooker
	WS     = 4 //washer
)

func main() {

	//dat := provider.Hatcher(AIRCON)
	//dat := provider.Hatcher(WS)
	//dat := provider.Hatcher(FRIDGE)
	//dat := provider.Hatcher(RCKR)
	//dat := provider.Hatcher(STOVE)

	//datBase := base64.StdEncoding.EncodeToString(dat)
	//d := []byte("!!" + datBase + "\n")
	send()
	//for ; ; {
	//
	//	flume()
	//	time.Sleep(5 * time.Second)
	//}

	//data := []byte{102, 103, 1, 0, 32, 89, 160, 181, 107, 48, 0, 0, 0, 0, 0, 0, 32, 89, 160, 181, 107, 48, 0, 1, 16, 2, 0, 0, 0, 1, 4, 0, 17, 3, 13, 13, 55, 46, 2, 123, 30, 190, 120, 12, 98, 173, 243, 187, 183, 239, 97, 115, 138, 120, 79, 90, 70, 13, 63, 185, 2, 114, 41, 9, 57, 2, 210, 214, 190, 185, 183, 247, 217, 209, 242, 62, 80, 232, 107, 9, 167, 118, 31, 179, 181, 192, 0, 235, 2, 140, 12, 44, 209, 237, 143, 101, 157, 4, 36, 250, 9, 152, 129, 53, 27, 153, 176, 187, 104, 36, 13, 76, 1, 196, 249, 218, 37, 254, 216141, 75, 40, 33, 27, 69, 146, 61, 194, 179, 19, 16, 174, 119, 160, 178, 161, 43, 215, 217, 185, 13, 190, 117, 42, 131, 74, 221, 66, 232, 157, 149, 63, 99, 143, 166, 229, 57, 160, 92, 214, 46, 140, 167, 205, 196, 72, 147, 230, 168, 45, 82, 254, 241, 184, 205, 26, 184, 26, 167, 132, 7, 188, 192, 46, 113, 44, 16, 178, 246, 238, 15, 195, 171, 183, 55, 2, 167, 199, 235, 93, 182, 198, 237, 93, 63, 206, 248, 157, 42, 25, 231, 30, 167, 141, 185, 106, 25, 70, 235252, 98, 169, 40, 150, 16, 103, 173, 252, 150, 236, 167, 253, 1, 236, 135, 60, 31, 163, 83, 35, 76, 178, 243, 96, 212, 203, 156, 145, 204, 78, 98, 26, 110, 16, 124, 45, 159, 118, 175, 135, 9, 58, 121, 73, 32, 16, 113, 216, 79, 229, 51, 247, 252, 147, 82, 75, 34, 120, 67, 62, 251, 215, 131, 184, 69, 30, 246, 112, 186, 20, 227, 26, 219, 15, 192}
	//dat := provider.Purifier("20110701")
	//fmt.Println(data)
	//send(data)
}

func flume() {
	conn := connector.Connection()
	fmt.Println("connection established!")

	dat := time.Now().UTC().Format("2006-01-02 15:04:05") + " INFO [com.gree.GenDevLog:40] asvr10.1.1.1 - devLogout," + macGen(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(64)) + "," + ipGen(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(64)) + "\n"
	//dat := time.Now().UTC().Format("2006-01-02 15:04:05") + " INFO [com.gree.GenDevLog:40] asvr10.1.1.1 - devLogin,f4911e000002,106.185.48.232" + "\n"
	//2015-10-06 16:27:44 INFO [com.gree.GenDevLog:40] asvr10.1.1.1 - devLogin,f4911e0bb55d,211.23.164.87
	fmt.Println(dat)
	conn.Write([]byte(dat))

}
func send() {

	conn := connector.Connection()
	fmt.Println("connection established!")

	//for i := 1; i < 2; i++ {
	//	dat := provider.Hatcher(AIRCON, macGen(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(64)))
	//	//dat := provider.Hatcher(WS)
		dat := provider.Hatcher(FRIDGE,macGen(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(16)))
	//	//dat := provider.Hatcher(RCKR)
	//	//dat := provider.Hatcher(STOVE)
	//	//dat := provider.Purifier("19880918",3)
		conn.Write(dat)
	//	log.Println(i)
	//	time.Sleep(700000 * time.Microsecond)
		log.Println(dat)
	//}

	//for i := 1; i < 1024; i++ {
	//	mac := macGen(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(8))
	//	airConCMD01 := provider.AirConCMD01(1, 1, 0, 1, 1, "18", 0, 0, 1, 0, 0, 0)
	//	log.Printf("开关机: %d, 模式: %d, 睡眠: %d, 扫风: %d, 风速: %d, 温度: %s, 灯光: %d, 超强: %d, 上下扫风: %d, 左右扫风: %d, 睡眠4: %d, 睡眠23: %d\n ", 1, 1, 0, 1, 1, "18", 0, 0, 1, 0, 0, 0)
	//	dat := provider.AirConTest(mac, airConCMD01)
	//	conn.Write(dat)
	//	time.Sleep(7000 * time.Microsecond)

	//	dat2 := provider.AirConTest(macGen(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(16)), provider.AirConCMD01(1, 2, 0, 1, 1, "25", 1, 0, 0, 1, 0, 1))
	//	log.Printf("开关机: %d, 模式: %d, 睡眠: %d, 扫风: %d, 风速: %d, 温度: %s, 灯光: %d, 超强: %d, 上下扫风: %d, 左右扫风: %d, 睡眠4: %d, 睡眠23: %d\n ", 1, 2, 0, 1, 1, "25", 1, 0, 0, 1, 0, 1)
	//	conn.Write(dat2)
	//	time.Sleep(70000 * time.Microsecond)
	//
	//	dat4 := provider.AirConTest(macGen(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(32)), provider.AirConCMD01(1, 3, 1, 0, 0, "16", 0, 1, 0, 0, 1, 1))
	//	log.Printf("开关机: %d, 模式: %d, 睡眠: %d, 扫风: %d, 风速: %d, 温度: %s, 灯光: %d, 超强: %d, 上下扫风: %d, 左右扫风: %d, 睡眠4: %d, 睡眠23: %d\n ", 1, 3, 1, 0, 0, "16", 0, 1, 0, 0, 1, 1)
	//	conn.Write(dat4)
	//	time.Sleep(700000 * time.Microsecond)
	//
	//	dat5 := provider.AirConTest(macGen(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(64)), provider.AirConCMD01(0, 0, 1, 0, 0, "25", 0, 1, 0, 0, 1, 1))
	//	log.Printf("开关机: %d, 模式: %d, 睡眠: %d, 扫风: %d, 风速: %d, 温度: %s, 灯光: %d, 超强: %d, 上下扫风: %d, 左右扫风: %d, 睡眠4: %d, 睡眠23: %d\n ", 0, 0, 1, 0, 0, "25", 0, 1, 0, 0, 1, 1)
	//	conn.Write(dat5)
	//	time.Sleep(70000 * time.Microsecond)
	//
	//	dat3 := provider.AirConTest(macGen(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(4)), provider.AirConCMD01(1, 1, 0, 1, 1, "20", 1, 0, 1, 0, 0, 1))
	//	log.Printf("开关机: %d, 模式: %d, 睡眠: %d, 扫风: %d, 风速: %d, 温度: %s, 灯光: %d, 超强: %d, 上下扫风: %d, 左右扫风: %d, 睡眠4: %d, 睡眠23: %d\n ", 1, 1, 0, 1, 1, "20", 1, 0, 1, 0, 0, 1)
	//	conn.Write(dat3)
	//	log.Println(i)
	//	time.Sleep(70000 * time.Microsecond)
	//}

	defer conn.Close()
}

func tmCus(date string) []byte {
	d := make([]byte, 3)

	Y, _ := strconv.Atoi(date[1:4])

	M, _ := strconv.Atoi(date[4:6])

	D, _ := strconv.Atoi(date[6:8])

	d[0] = byte(Y)

	d[1] = byte(M)

	d[2] = byte(D)

	return d
}

func macGen(i int) string {

	macs := []string{"2059a0b50079", "2059a0b56a44", "845dd74f2c11", "f4911e0bb55d", "f4911e0bb545", "f4911e10afef", "845dd74f2c12", "f4911e10afb7", "2059a0b40d0f", "2059a0b44aad", "f4911e10af8b", "f4911e1562af", "34ea3400ccf6", "2059a0b44aa6", "2059a0b44aa5", "f4911e15629f", "f4911e0bb555", "f4911e096b5b", "f4911e0bc235", "f4911e0bb567", "f4911e0bb56a", "f4911e0bb56b", "f4911e1562c1", "f4911e156293", "f4911e1562c3", "f4911e156297", "f4911e1562a0", "f4911e1562b4", "f4911e156286", "f4911e0bb54e", "2059a0b56b30", "2059a0b56b2a", "2059a0b56ae5", "2059a0b56b1d", "2059a0b4329b", "2059a0b56af3", "2059a0b56af2", "2059a0b56ada", "2059a0b56ad4", "f4911e10afc9", "2059a0b56b10", "f4911e096000", "f4911e095ebb", "f4911e095b57", "f4911e095e60", "f4911e1562a2", "313132323333", "112233445566", "f4911e1913dc", "112233445567", "f4911e1059ba", "f4911e054cd8", "f4911e048652", "f4911e0ccfcb", "f4911e1ec04b", "f4911e1e3fa0", "f4911e1e63ba", "f4911e1059b5", "f4911e1ef328", "f4911e12a0f2", "f4911e1e6e59", "f4911e1e5399", "f4911e1f0619", "f4911e10b03c", "f4911e21ca21", "f4911e12a619", "f4911e12b4c5", "f4911e1f0849", "f4911e1fbdb3", "f4911e272134", "f4911e271aff", "f4911e12bc02", "f4911e271f63", "f4911e12a87e", "f4911e20f48d", "f4911e23eb79", "2059a0b56ad0", "f4911e12bea2", "f4911e12c136", "f4911e1e5c87", "f4911e12c926", "f4911e12c5a9"}

	return macs[i]

}

func ipGen(i int) string {
	ips := []string{"211.23.164.87",
			"103.250.13.140",
			"182.163.234.137",
			"106.185.48.232",
			"46.21.157.23",
			"116.6.118.62",
			"218.7.9.1",
			"218.7.9.2",
			"218.7.9.3",
			"218.7.9.8",
			"218.7.9.9",
			"218.7.9.10",
			"218.7.9.11",
			"218.7.90.1",
			"218.7.90.2",
			"218.7.90.3",
			"61.143.58.0",
			"61.143.58.1",
			"61.143.58.2",
			"61.143.58.3",
			"61.143.58.4",
			"61.143.58.5",
			"61.143.58.6",
			"61.143.58.7",
			"61.143.58.8",
			"61.143.58.9",
			"61.143.58.10",
			"61.143.58.11",
			"61.143.58.12",
			"61.143.58.13",
			"61.143.58.14",
			"61.143.58.15",
			"61.143.58.16",
			"61.143.58.17",
			"61.143.58.18",
			"61.143.58.19",
			"61.143.58.20",
			"61.143.58.21",
			"61.143.58.22",
			"61.143.58.23",
			"61.143.58.24",
			"61.143.59.1",
			"61.143.59.2",
			"61.143.59.3",
			"61.143.59.4",
			"61.143.59.5",
			"61.143.59.6",
			"61.143.59.7",
			"61.143.59.8",
			"61.143.59.9",
			"61.143.59.10",
			"61.143.59.11",
			"61.143.59.12",
			"61.143.59.13",
			"61.143.59.14",
			"61.143.59.15",
			"61.143.59.16",
			"61.143.59.17",
			"61.143.59.18",
			"61.143.59.19",
			"61.143.59.20",
			"61.143.59.21",
			"61.143.59.22",
			"61.143.59.23",
			"61.143.59.24"}
	return ips[i]
}
