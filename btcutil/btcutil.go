package btcutil

import (
	"bytes"
	"errors"
	"github.com/hoonrepo/goUtils/hexutil"
	"math/big"
)




var (
	//
	maxTarget_1_Bits = big.NewInt(0x1D00ffff)
	//0x00000000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF
	maxTarget_1_Difficult = new(big.Int).Exp(big.NewInt(2),big.NewInt(224),big.NewInt(0))
)

//eg.
//0X1729D72D
//Target Threshold = 0x29D72D * 256** (0x17 - 3))
//bits: 8bytes/64bits
func CalculateTargetThresholdFromBits(bits *big.Int) (targetThreshold *big.Int){

	calConst := big.NewInt(0x00FFFFFF)
	//Fetch 0xYY by 0xAABBBBBB >> 24
	exponent := new(big.Int).Rsh(bits,24)
	//Fetch 0xBBBBBB by 0xAABBBBBB & 0x00ffffff
	significand := new(big.Int).And(bits,calConst)
	//0xBBBBBB * 256 ** (0xAA - 3)
	multiplier := new(big.Int).Exp(
		big.NewInt(256),
		new(big.Int).Sub(exponent,big.NewInt(3)),
		big.NewInt(0))


	targetThreshold = new(big.Int).Mul(significand,multiplier)
	return
}

//0x00000000FFFF0000000000000000000000000000000000000000000000000000
func CalculateMaxTargetThreshold()(targetThreshold *big.Int){

	targetThreshold = CalculateTargetThresholdFromBits(maxTarget_1_Bits)
	return
}

// pdiff = 0x00000000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF /
// currentDiff
func GetPdiff (bits *big.Int)(pdiff *big.Float){

	//0x00000000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF

	maxTagDiff_1 := new(big.Float).SetInt(maxTarget_1_Difficult)
	currentDiff := new(big.Float).SetInt(CalculateTargetThresholdFromBits(bits))
	pdiff = new(big.Float).Quo(maxTagDiff_1,currentDiff)
	return

}

// bdiff = 0x00000000FFFF0000000000000000000000000000000000000000000000000000 /
// currentDiff
func GetBdiff(bits *big.Int)(bdiff *big.Float){

	//The highest possible target (difficulty 1)
	// is defined as 0x1d00ffff

	maxTagDiff_1 := new(big.Float).SetInt(CalculateMaxTargetThreshold())
	currentDiff := new(big.Float).SetInt(CalculateTargetThresholdFromBits(bits))
	bdiff = new(big.Float).Quo(maxTagDiff_1,currentDiff)

	return
}

func ToLittleEndian(org interface{})(string,error) {

	var b []byte
	switch o := org.(type) {

	case []byte:
		b = []byte(o)
	case string:
		bb,err := hexutil.Decode(string(o))
		if err != nil {
			return "nil string",err
		}
		b = bb
	default:
		return "nil string",errors.New("The type is not supported")

	}

	//buffer := &bytes.Buffer{} //bytes.NewBuffer([]byte{})
	//binary.Write(buffer,binary.LittleEndian,b)

	buffer := &bytes.Buffer{}
	//result := &[32]byte{}

	for i := len(b) - 1 ; i >= 0 ; i--{
		buffer.WriteByte(b[i])
	}

	return hexutil.Encode(buffer.Bytes()),nil
}