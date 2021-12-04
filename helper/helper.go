package helper

import "strconv"

func SliceUint64ToString(s []uint64) []string {
	var ss []string
	for _, v := range s {
		ss = append(ss, strconv.FormatUint(v, 10))
	}

	return ss
}
