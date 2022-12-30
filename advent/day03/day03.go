package day03

import "fmt"

func Priority(ch rune) (int, error) {
	ascii := int(ch)

	switch {
	case ascii > 64 && ascii < 91:
		return ascii - 64 + 26, nil
	case ascii > 96 && ascii < 123:
		return ascii - 96, nil
	default:
		return -1, fmt.Errorf("not an alpha character: %v", ch)
	}

}

func Split(bag string) (string, string) {
	return bag[:len(bag)/2], bag[len(bag)/2:]
}
