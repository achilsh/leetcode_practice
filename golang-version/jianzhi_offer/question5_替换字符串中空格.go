package jianzhi_offer

func ReplaceSpaceInStr(data string) string {
	if len(data) <= 0 {
		return data
	}
	var spaceNums, totalNums = 0, 0
	for i := 0; i < len(data); i++ {
		if data[i] == ' ' {
			spaceNums++
		} else {
			totalNums++
		}
	}
	newtotalNums := totalNums + spaceNums*3

	var dst []rune = make([]rune, newtotalNums)
	for i, j := len(data)-1, newtotalNums-1; i >= 0; i-- {
		if data[i] != ' ' {
			dst[j] = rune(data[i])
			j--
			continue
		}
		dst[j] = '0'
		j--

		dst[j] = '2'
		j--

		dst[j] = '%'
		j--
	}
	return string(dst)
}
