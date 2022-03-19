package fullymal94

import "math"

func Convert(fullymalNumber string) int {

	letters := []string{}

	for i := 33; i < 127; i++ {
		letters = append(letters, string(i))
	}

	indexNumbers := []int{}

	numberBase10 := 0

	for _, v := range fullymalNumber {

		indexNumbers = append(indexNumbers, getElementIndex(letters, string(v)))

	}

	expCount := len(indexNumbers) - 1

	for c := 0; c < len(indexNumbers)-1; c++ {

		numberBase10 += int(math.Pow(float64(92), float64(expCount)) * float64(indexNumbers[c]))

		expCount--

	}

	numberBase10 += indexNumbers[len(indexNumbers)-1]

	return numberBase10

}

func getElementIndex(letters []string, val string) int {
	for i, v := range letters {
		if val == v {
			return i
		}
	}

	return -1
}
