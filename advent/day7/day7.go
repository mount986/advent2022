package day7

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

func BuildFileTree(scanner *bufio.Scanner) error {
	var cDir Dir

	for scanner.Scan() {
		line := scanner.Text()

		re, err := regexp.Compile(`(\$)? ?(\w+) ?(\S*)`)
		if err != nil {
			return nil
		}

		matches := re.FindStringSubmatch(line)

		if matches[1] == "$" {
			switch matches[2] {
			case "cd":
				cDir = cd(cDir, matches[3])
			case "ls":

			default:
				return fmt.Errorf("invalid command: %v", matches[2])
			}
		} else {
			if matches[2] != "dir" {
				if !cDir.ContainsFile(matches[3]) {
					size, err := strconv.Atoi(matches[2])
					if err != nil {
						return err
					}
					cDir.NewFile(matches[3], size)
				}
			}
		}
	}

	return nil
}

func cd(dir Dir, sub string) Dir {
	if sub == "/" {
		return ROOT
	}

	if sub == ".." {
		return *dir.Parent
	}

	if dir.ContainsSub(sub) {
		return dir.Subs[sub]
	}

	return dir.NewSub(sub)
}

func CollectSmall(dir Dir, max int) ([]Dir, int) {
	var small []Dir
	sum := 0

	if dir.Size() < max {
		small = append(small, dir)
		sum += dir.Size()
	}

	for _, sub := range dir.Subs {
		sm, su := CollectSmall(sub, max)
		small = append(small, sm...)
		sum += su
	}

	return small, sum
}

func CollectLarge(dir Dir, min int) []Dir {
	var largeDirs []Dir

	if dir.Size() >= min {
		largeDirs = append(largeDirs, dir)

		for _, sub := range dir.Subs {
			largeDirs = append(largeDirs, CollectLarge(sub, min)...)
		}
	}

	return largeDirs
}
