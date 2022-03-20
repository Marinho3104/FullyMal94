package fullymal94

import (
	"math"
)

func setFullyMalNumber(number int) string {

	fullymalNumber := ""

	letters := []string{}

	for i := 33; i < 127; i++ {
		letters = append(letters, string(i))
	}

	////////

	for {

		if number <= 94 {
			break
		}

		intPart := math.Floor(float64(number) / float64(94))

		val := number - (int(intPart) * 94)

		fullymalNumber = letters[val] + fullymalNumber

		number = int(intPart)

	}

	if number != 0 {
		fullymalNumber = letters[number] + fullymalNumber
	}

	return fullymalNumber

}
