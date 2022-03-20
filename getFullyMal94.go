package fullymal94

import (
	"math/big"
)

func setFullyMalNumber(number *big.Int) string {

	fullymalNumber := ""

	letters := []string{}

	for i := 33; i < 127; i++ {
		letters = append(letters, string(i))
	}

	////////

	for {

		if number.Cmp(big.NewInt(95)) == -1 {
			break
		}

		intPart := new(big.Int)

		intPart.Div(number, big.NewInt(94))

		mutlIntPartTo94 := new(big.Int)

		mutlIntPartTo94.Mul(intPart, big.NewInt(94))

		val := new(big.Int)

		val.Sub(number, mutlIntPartTo94)

		fullymalNumber = letters[val.Int64()] + fullymalNumber

		number = intPart

	}

	if number.Cmp(big.NewInt(0)) != 0 {
		fullymalNumber = letters[number.Int64()] + fullymalNumber
	}

	return fullymalNumber

}
