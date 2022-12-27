package advent

import (
	"bytes"
	"io"
)

func (a *Advent) Day6Part1() (int, error) {
	buf := bytes.NewBuffer(nil)

	io.Copy(buf, a.Input)
	input := buf.Bytes()	

	return findUnique(input, 4)
}

func (a *Advent) Day6Part2() (int, error) {
	buf := bytes.NewBuffer(nil)

	io.Copy(buf, a.Input)
	input := buf.Bytes()	

	return findUnique(input, 14)
}

func unique(intSlice []byte) []byte {
    keys := make(map[byte]bool)
    list := []byte{}	
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}

func findUnique(input []byte, buff int) (int, error) {
	

	for i := 0; i < len(input)-buff; i++ {
		sub := input[i : i+buff]
		if len(unique(sub)) == buff {
			return i + buff, nil
		}
	}

	return 0, nil
}