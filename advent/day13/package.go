package day13

import "bufio"

type Package struct {
	lList List
	rList List
}

func NewPackage(scanner *bufio.Scanner) (Package, error) {
	var p Package
	var err error

	p.lList, err = BuildList(scanner)
	if err != nil {
		return p, err
	}
	scanner.Scan()
	scanner.Scan()

	p.rList, err = BuildList(scanner)
	if err != nil {
		return p, err
	}
	scanner.Scan()
	scanner.Scan()

	return p, nil
}

func (p *Package) ValidPackage() bool {
	return p.lList.Compare(p.rList) <= 0
}
