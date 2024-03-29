package elementary

// IsAnagram compares two strings by sorting them and comparing
func IsAnagram(word1, word2 string) bool {
	byte1 := []byte(word1)
	byte2 := []byte(word2)
	quickSort(byte1, 0, len(byte1)-1)
	quickSort(byte2, 0, len(byte2)-1)
	return string(byte1) == string(byte2)
}

// IsAnagramFast compares two strings by counting chars in a map
func IsAnagramFast(word1, word2 string) bool {
	byte1 := []byte(word1)
	byte2 := []byte(word2)

	map1 := make(map[string]int)
	for i := 0; i < len(byte1); i++ {
		if val, ok := map1[string(byte1[i])]; ok {
			map1[string(byte1[i])] = val + 1
		} else {
			map1[string(byte1[i])] = 1
		}
	}

	for i := 0; i < len(byte2); i++ {
		if val, ok := map1[string(byte2[i])]; ok {
			map1[string(byte2[i])] = val - 1
			if map1[string(byte2[i])] < 0 {
				return false
			}
		} else {
			return false
		}
	}

	for _, element := range map1 {
		if element != 0 {
			return false
		}
	}

	//fmt.Println("map:", map1)
	return true
}

// copied from wikipedia from pseudo-code
func quickSort(slice []byte, low, high int) {
	if low < high {
		p := quickPartition(slice, low, high)
		quickSort(slice, low, p-1)
		quickSort(slice, p+1, high)
	}
}

func quickPartition(slice []byte, low, high int) int {
	pivot := slice[high]
	i := low
	var buf byte
	for j := low; j <= high; j++ {
		if slice[j] < pivot {
			buf = slice[j]
			slice[j] = slice[i]
			slice[i] = buf
			i++
		}
	}
	buf = slice[high]
	slice[high] = slice[i]
	slice[i] = buf
	return i
}
