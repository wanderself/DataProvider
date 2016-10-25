package provider

func Hatcher(code int) []byte {
	dat := []byte{123, 124, 0, 0, 1, 40, 0, 1, 0, 1, 1, 40, 0, 1, 0, 11, 40, 0, 1, 0, 11, 40, 0, 1, 0, 11, 40, 0, 1, 0, 11, 40, 0, 1, 0, 11, 40, 0, 1, 0, 11, 40, 0, 1, 0, 11, 40, 0, 1, 0, 11, 40, 0, 1, 0, 11, 40, 0, 1, 0, 1}
	//handle.Log("original:", dat)

	stvMid := []byte{0, 0, 134, 0}
	acMid := []byte{0, 1, 0, 1}

	switch code {
	case 0:

		copy(dat[22:26], stvMid)
		copy(dat[30], []byte{1})
		return dat
	case 1:
		copy(dat[22:26], stvMid)

		copy(dat[30], []byte{49})
		return dat
	case 2:

		copy(dat[22:26], stvMid)
		copy(dat[30], []byte{52})
		return dat
	case 3:
		copy(dat[22:26], acMid)
		copy(dat[30], []byte{1})
		return dat
	case 4:
		copy(dat[22:26], acMid)
		copy(dat[30], []byte{1})
		return dat

	}

	return dat
}
