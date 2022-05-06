package fullymal94

import (
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

func (fullyMal *FullyMal) CheckTable() {
	if len(fullyMal.FullyMalTable) != 94 {
		fullyMal.FullyMalTable = []string{}

		fullyMal.SetTableValue()
	}
}

func (fullyMal *FullyMal) FromBytes(_bytes []byte) {

	fullyMal.CheckTable()

	if len(_bytes) == 0 {
		fullyMal.FullyMalRepres = ""
		return
	}

	valueBigInt := new(big.Int)

	valueBigInt.SetBytes(_bytes)

	fullyMal.FromInt(valueBigInt)

}

func (fullyMal *FullyMal) ToBytes() []byte {

	fullyMal.CheckTable()

	if fullyMal.FullyMalRepres == "" {
		return []byte("")
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
			return []byte("")
		}

		byteArray = append(byteArray, byte(num))
	}

	return byteArray

}

func (fullymal *FullyMal) ToInt() *big.Int {

	posValue := big.NewInt(1)

	fullymal.CheckTable()

	if fullymal.FullyMalRepres[:2] == "- " {
		posValue = big.NewInt(-1)
		fullymal.FullyMalRepres = fullymal.FullyMalRepres[2:]
	}

	if fullymal.FullyMalRepres == "" {
		return big.NewInt(0)
	}

	indexNumbers := []int{}

	numberBase10 := new(big.Int)

	for _, v := range fullymal.FullyMalRepres {

		elementIndex := getElementIndex(fullymal.FullyMalTable, string(v))

		if elementIndex == -1 {
			return big.NewInt(0)
		}

		indexNumbers = append(indexNumbers, elementIndex)

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

	numberBase10.Mul(numberBase10, posValue)

	return numberBase10

}

func (fullyMal *FullyMal) ToBinary() string {

	fullyMal.CheckTable()

	binaryValue := ""

	decimalNumber := fullyMal.ToInt()

	posValue := ""

	if decimalNumber.Cmp(big.NewInt(0)) == -1 {
		decimalNumber.Mul(decimalNumber, big.NewInt(-1))
		posValue = "-"
	}

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

	binaryValueCorrect = posValue + binaryValueCorrect

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

	posValue := ""

	if number.Cmp(big.NewInt(0)) == -1 {
		posValue = "- "

		number.Mul(number, big.NewInt(-1))
	}

	fullyMal.CheckTable()

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

	fullyMal.FullyMalRepres = posValue + fullyMal.FullyMalRepres

}
