package advent

func uniqueBytes(byteSlice []byte) []byte {
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

func uniqueStrings(stringSlice []string) []string {
    keys := make(map[string]bool)
    list := []string{}	
    for _, entry := range stringSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}