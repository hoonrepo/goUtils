package btcutil

import (
	"crypto/sha256"
	"fmt"
	"github.com/hoonrepo/goUtils/hexutil"
	"math/big"
	"strings"
)

type Block struct {
	number     	uint
	nonce      	*big.Int
	bits       	*big.Int
	timestamp  	*big.Int
	version		*big.Int
	merkleRoot 	*big.Int
	prevHash   	*big.Int
}

func NewBlock(number uint,
	nonce int64,
	bits int64,
	timestamp int64,
	version int64,
	merkleRoot string,
	prevHash string) *Block {

	block := &Block{}
	block.number = number
	block.nonce = big.NewInt(nonce)
	block.bits = big.NewInt(bits)
	block.timestamp = big.NewInt(timestamp)
	block.version = big.NewInt(version)
	mr,err := hexutil.Decode256(merkleRoot)
	if err != nil {
		panic(err)
	}
	block.merkleRoot = new(big.Int).SetBytes(mr)
	ph,errPh := hexutil.Decode256(prevHash)
	if errPh != nil {
		panic(errPh)
	}
	block.prevHash = new(big.Int).SetBytes(ph)


	return block

}

func (b *Block) GetNumber() uint{
	return b.number
}

func (b *Block) GetNonce() uint64{
	return b.nonce.Uint64()
}

func (b *Block) GetBits() uint64{
	return b.bits.Uint64()
}

func (b *Block)GetTimeStamp() int64{

	return b.timestamp.Int64()
}

func (b *Block)GetVersion() string{
	return hexutil.EncodeBig(b.version)
}

func (b *Block) GetMerkleRoot() string{
	return hexutil.EncodeBig(b.merkleRoot)
}

func (b *Block) GetPrevHash() string{
	return hexutil.EncodeBig(b.prevHash)
}



/*
	https://en.bitcoin.it/wiki/Block_hashing_algorithm
	$header_hex = $version . $prevBlockHash . $rootHash . $time . $bits . $nonce;


	125552

	>>> import hashlib
	>>> header_hex = ("01000000" +
	 "81cd02ab7e569e8bcd9317e2fe99f2de44d49ab2b8851ba4a308000000000000" +
	 "e320b6c2fffc8d750423db8b1eb942ae710e951ed797f7affc8892b0f1fc122b" +
	//2b12fcf1b09288fcaff797d71e950e71ae42b91e8bdb2304758dfcffc2b620e3
	 "c7f5d74d" +
	 "f2b9441a" +  //1a44b9f2
	 "42a14695")	//9546a142
	>>> header_bin = header_hex.decode('hex')
	>>> hash = hashlib.sha256(hashlib.sha256(header_bin).digest()).digest()
	>>> hash.encode('hex_codec')
	'1dbd981fe6985776b644b173a4d0385ddc1aa2a829688d1e0000000000000000'
	>>> hash[::-1].encode('hex_codec')
	'00000000000000001e8d6829a8a21adc5d38d0a473b144b6765798e61f98bd1d'
	 */
func (b *Block) POW() (blockHash string,err error){

	blockHash = "nil string"

	version := hexutil.Encode(b.version.Bytes())
	preHash,phErr := hexutil.Encode256(b.prevHash.Bytes())
	if phErr != nil {
		err = phErr
		return
	}
	merkleRoot,mrErr := hexutil.Encode256(b.merkleRoot.Bytes())
	if mrErr != nil {
		err = mrErr
		return
	}
	timestamp := hexutil.Encode(b.timestamp.Bytes())
	bits := hexutil.Encode(b.bits.Bytes())
	nonce := hexutil.Encode(b.nonce.Bytes())

	versionLittleEndian,vErr := ToLittleEndian(version)
	if vErr != nil {
		err = vErr
		return
	}
	preHashLittleEndian,pErr := ToLittleEndian(preHash)
	if pErr != nil {
		err = pErr
		return
	}
	merkleRootLittleEndian,mErr := ToLittleEndian(merkleRoot)
	if mErr != nil {
		err = mErr
		return
	}
	timestampEndian,tErr := ToLittleEndian(timestamp)
	if tErr != nil {
		err = tErr
		return
	}
	bitsLittleEndian,bErr := ToLittleEndian(bits)
	if bErr != nil {
		err = bErr
		return
	}
	nonceRootLittleEndian,nErr := ToLittleEndian(nonce)
	if nErr != nil {
		err = nErr
		return
	}

	sb := &strings.Builder{}
	sb.WriteString("0x")
	sb.WriteString(versionLittleEndian[2:])
	sb.WriteString(preHashLittleEndian[2:])
	sb.WriteString(merkleRootLittleEndian[2:])
	sb.WriteString(timestampEndian[2:])
	sb.WriteString(bitsLittleEndian[2:])
	sb.WriteString(nonceRootLittleEndian[2:])

	org := sb.String()
	fmt.Printf("strsum= %s\n",org)

	totalHex,err := hexutil.Decode(org)
	if err != nil {
		panic(err)
	}

	hash1st := sha256.New()
	hash1st.Write(totalHex)
	sha256Result1 := hash1st.Sum(nil)

	hash2nd := sha256.New()
	hash2nd.Write(sha256Result1)
	powHashResult := hash2nd.Sum(nil)

	powHashStr := hexutil.Encode(powHashResult[:])
	r,covErr := ToLittleEndian(powHashStr)
	if covErr != nil {
		err = covErr
		return
	}
	blockHash = r
	return

}