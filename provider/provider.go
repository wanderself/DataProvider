package provider

import "fmt"

func Hatcher() {
	dat := []byte{123, 124, 0, 0, 1, 40, 0, 1, 0, 1}

	fmt.Println(dat)

	mid := []byte{0, 0, 134, 0}

	copy(dat[4:8], mid)

	fmt.Println(dat)
}
