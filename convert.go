package fullymal94

import (
	"math/big"
)

func Convert(fullymalNumber string) *big.Int {

	letters := []string{}

	for i := 33; i < 127; i++ {
		letters = append(letters, string(i))
	}

	indexNumbers := []int{}

	numberBase10 := new(big.Int)

	for _, v := range fullymalNumber {

		indexNumbers = append(indexNumbers, getElementIndex(letters, string(v)))

	}

	expCount := len(indexNumbers) - 1

	for c := 0; c < len(indexNumbers)-1; c++ {

		_pow := new(big.Int)

		_pow.Exp(big.NewInt(94), big.NewInt(int64(expCount)), nil)

		_mul := new(big.Int)

		_mul.Mul(_pow, big.NewInt(int64(indexNumbers[c])))

		numberBase10.Add(numberBase10, _mul)

		expCount--

	}

	numberBase10.Add(numberBase10, big.NewInt(int64(indexNumbers[len(indexNumbers)-1])))

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
