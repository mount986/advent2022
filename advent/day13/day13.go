package day13

import "bufio"

func BuildPackages(scanner *bufio.Scanner) ([]Package, error) {
	var pList []Package

	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		p, err := NewPackage(scanner)
		if err != nil {
			return pList, err
		}
		pList = append(pList, p)
	}

	return pList, nil
}
