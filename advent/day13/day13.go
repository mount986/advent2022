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

func CombineLists(pList []Package) []List {
	var lists []List

	for _, p := range pList {
		lists = append(lists, p.lList)
		lists = append(lists, p.rList)
	}

	return lists
}