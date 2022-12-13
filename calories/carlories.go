package calories

import (
	"bufio"
	"os"
	"strconv"
)

func CaloriesMax(input *os.File) (int, error) {
	max := 0
	sum := 0

	scanner := bufio.NewScanner(input)
    for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sum = 0
		} else {
			val, err := strconv.Atoi(line)
			if err != nil {
				return 0, err
			}

			sum += val
			if sum > max {
				max = sum
			}
		}
    }

	return max, nil
}

func CaloriesMax3(input *os.File) (int, error) {
	max1, max2, max3, sum := 0, 0, 0, 0

	scanner := bufio.NewScanner(input)
    for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if sum > max1 {
				max3 = max2
				max2 = max1
				max1 = sum				
			} else if sum > max2 {
				max3 = max2
				max2 = sum
			} else if sum > max3 {
				max3 = sum
			}

			sum = 0
		} else {
			val, err := strconv.Atoi(line)
			if err != nil {
				return 0, err
			}

			sum += val
		}
    }

	if sum > max1 {
		max3 = max2
		max2 = max1
		max1 = sum				
	} else if sum > max2 {
		max3 = max2
		max2 = sum
	} else if sum > max3 {
		max3 = sum
	}

	return max1 + max2 + max3, nil
}