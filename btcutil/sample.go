package btcutil

import (
	"fmt"
	"github.com/hoonrepo/goUtils/hexutil"
	"gopkg.in/blockcypher/gobcy.v1"
	"math/big"
	"time"
)

/*
Number	539753
Hash	000000000000000000031d0abc975b4b049bc245e123432b3e3d2b129a21701e
Timestamp	2018-09-03 07:05:29 / 2018-09-03 15:05:29(Local)
Version	0x20000000
Bits	388618029/0x1729d72d
Nonce 	105721929/0x064d3049
Merkle Root ff6b6d3a43a1c870db54d2ce59b5f6cbde704d8eee02ee1bf37c7aac36acda0c
prev_hash	000000000000000000205e3cdab18cea24d526b760a864d0f5a7c9ec0d5aea55

 */
func Sample539753() (block *Block){


	ts := time.Date(2018,9,3,15,5,29,0,time.Local).Unix()

	block = NewBlock(
		539753,
		0x064d3049,
		0x1729d72d,
		ts,
		0x20000000,
		"0xff6b6d3a43a1c870db54d2ce59b5f6cbde704d8eee02ee1bf37c7aac36acda0c",
		"0x000000000000000000205e3cdab18cea24d526b760a864d0f5a7c9ec0d5aea55",
		)

	return
}

func Check539753()(bool){
	block := Sample539753()
	hashStr,e := block.POW()
	if e != nil {
		fmt.Printf("error= %s",e.Error())
		return false
	}

	hashBytes,eDe := hexutil.Decode256(hashStr)
	if eDe != nil {
		fmt.Printf("error= %s",eDe.Error())
		return false
	}

	hashBig := new(big.Int).SetBytes(hashBytes)

	target := CalculateTargetThresholdFromBits(block.bits)

	hr,_ := hexutil.Encode256(hashBytes)
	tr,_ := hexutil.Encode256(target.Bytes())
	fmt.Printf("hashBig= %s\n",hr)
	fmt.Printf("target= %s\n",tr)

	if hashBig.Cmp(target) <= 0 {
		return true
	}else{
		return false
	}
}

func buildUtilBlock(block gobcy.Block)*Block{
	utilBlock := NewBlock(
		uint(block.Height),
		int64(block.Nonce),
		int64(block.Bits),
		block.Time.Unix(),
		int64(block.Ver),
		fmt.Sprintf("0x%s",block.MerkleRoot),
		fmt.Sprintf("0x%s",block.PrevBlock),
	)

	return utilBlock
}

func CheckPow4RealBlock(block gobcy.Block)(bool){

	blockHash := fmt.Sprintf("0x%s",block.Hash)
	fmt.Printf("blockHash= %s\n",blockHash)

	utilBlock := buildUtilBlock(block)

	hash,err := utilBlock.POW()
	fmt.Printf("hash= %s\n",hash)

	if err != nil {
		panic(err.Error())
	}

	if blockHash == hash {

		return true
	}

	return false
}


func CheckCmp(block gobcy.Block)bool{

	blockHash := block.Hash
	fmt.Printf("blockHash= %s\n",blockHash)

	utilBlock := buildUtilBlock(block)
	powHashStr,e := utilBlock.POW()
	if e != nil {
		fmt.Printf("utilBlock.POW()= %s",e.Error())
		return false
	}

	powData,eD256 := hexutil.Decode256(powHashStr)
	if eD256 != nil {
		fmt.Printf("hexutil.Decode256= %s",eD256.Error())
		return false
	}

	hashBig := new(big.Int).SetBytes(powData)

	target := CalculateTargetThresholdFromBits(utilBlock.bits)

	hr,_ := hexutil.Encode256(powData)
	tr,_ := hexutil.Encode256(target.Bytes())
	fmt.Printf("hashBig= %s\n",hr)
	fmt.Printf("target= %s\n",tr)

	if hashBig.Cmp(target) <= 0 {
		return true
	}else{
		return false
	}
}