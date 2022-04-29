package fullymal94

import (
	"fmt"
	"math/big"
	"strconv"
)

type FullyMal struct {
	FullyMalTable  []string
	FullyMalRepres string
}

func (fullyMal *FullyMal) SetTableValue() {
	for i := 33; i < 127; i++ {
		fullyMal.FullyMalTable = append(fullyMal.FullyMalTable, string(i))
	}
}

func (fullyMal *FullyMal) FromBytes(_bytes []byte) {

	if len(fullyMal.FullyMalTable) != 94 {
		fullyMal.FullyMalTable = []string{}

		fullyMal.SetTableValue()
	}

	valueBigInt := new(big.Int)

	valueBigInt.SetBytes(_bytes)

	fullyMal.FromInt(valueBigInt)

}

func (fullyMal *FullyMal) ToBytes() []byte {

	if len(fullyMal.FullyMalTable) != 94 {
		fullyMal.FullyMalTable = []string{}

		fullyMal.SetTableValue()
	}

	byteArray := []byte{}

	binaryValueCorrect := fullyMal.ToBinary()

	for {
		if len(binaryValueCorrect)%8 == 0 {
			break
		}

		binaryValueCorrect = "0" + binaryValueCorrect
	}

	for i := 0; i <= len(binaryValueCorrect)-8; i += 8 {

		num, err := strconv.ParseUint(binaryValueCorrect[i:i+8], 2, 8)

		if err != nil {
			fmt.Println(err)
			return []byte("")
		}

		byteArray = append(byteArray, byte(num))
	}

	return byteArray

}

func (fullymal *FullyMal) ToInt() *big.Int {

	indexNumbers := []int{}

	numberBase10 := new(big.Int)

	for _, v := range fullymal.FullyMalRepres {

		indexNumbers = append(indexNumbers, getElementIndex(fullymal.FullyMalTable, string(v)))

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

func (fullyMal *FullyMal) ToBinary() string {

	binaryValue := ""

	decimalNumber := fullyMal.ToInt()

	for {

		intPart := new(big.Int)

		modPart := new(big.Int)

		intPart.Div(decimalNumber, big.NewInt(2))

		modPart.Mod(decimalNumber, big.NewInt(2))

		decimalNumber = intPart

		binaryValue += modPart.String()

		if intPart.Cmp(big.NewInt(0)) == 0 {
			break
		}

	}

	binaryValueCorrect := ""

	for i := len(binaryValue) - 1; i >= 0; i-- {
		binaryValueCorrect += string(binaryValue[i])
	}

	return binaryValueCorrect
}

func getElementIndex(letters []string, val string) int {
	for i, v := range letters {
		if val == v {
			return i
		}
	}

	return -1
}

func (fullyMal *FullyMal) FromInt(number *big.Int) {

	if len(fullyMal.FullyMalTable) != 94 {
		fullyMal.FullyMalTable = []string{}

		fullyMal.SetTableValue()
	}

	fullyMal.FullyMalRepres = ""

	for {

		if number.Cmp(big.NewInt(94)) == -1 {
			break
		}

		intPart := new(big.Int)

		intPart.Div(number, big.NewInt(94))

		mutlIntPartTo94 := new(big.Int)

		mutlIntPartTo94.Mul(intPart, big.NewInt(94))

		val := new(big.Int)

		val.Sub(number, mutlIntPartTo94)

		fullyMal.FullyMalRepres = fullyMal.FullyMalTable[val.Int64()] + fullyMal.FullyMalRepres

		number = intPart

	}

	if number.Cmp(big.NewInt(0)) != 0 {
		fullyMal.FullyMalRepres = fullyMal.FullyMalTable[number.Int64()] + fullyMal.FullyMalRepres
	}

}
