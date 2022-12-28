package day6

func Unique(byteSlice []byte) []byte {
    keys := make(map[byte]bool)
    list := []byte{}	
    for _, entry := range byteSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}

func FindUnique(input []byte, buff int) (int, error) {
	for i := 0; i < len(input)-buff; i++ {
		sub := input[i : i+buff]
		if len(Unique(sub)) == buff {
			return i + buff, nil
		}
	}

	return 0, nil
}