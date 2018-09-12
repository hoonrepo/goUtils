package btcutil

import (
	"flag"
	"fmt"
	"github.com/hoonrepo/goUtils/btcutil/cypherbtc"
	"testing"
)

var(
	token *string
)

func init() {
	token = flag.String("token","nil string","blockCypher's token")
	fmt.Printf("token= %s\n",*token)
}

func TestSample539753(t *testing.T) {

	block := Sample539753()
	r,e := block.POW()
	if e == nil {
		if r == "0x000000000000000000031d0abc975b4b049bc245e123432b3e3d2b129a21701e"{
			t.Logf("539753_pow= %s,Succeed!",r)
		}else{
			t.Logf("539753_pow= %s,Faild!",r)
		}

	}else{
		t.Errorf("error= %s",e.Error())
	}

}


func TestCheck539753(t *testing.T) {
	if Check539753() {
		t.Log("Succeed!")
	}else{
		t.Log("Failed!")
	}
}

func TestCheckPow4RealBlockLast10(t *testing.T) {

	var lastBlockHeight int

	btcApi := cypherbtc.NewBtcApi(*token)

	chain,err := btcApi.FetchChain()

	if err != nil {
		t.Errorf("FetchChain()_error= %s\n",err.Error())
		return
	}
	lastBlockHeight = chain.Height

	for i := lastBlockHeight ; i > lastBlockHeight - 10 ; i-- {
		t.Logf("block_height= %d",i)
		block,eB := btcApi.FetchBlock(i)
		if eB != nil {
			t.Errorf("FetchBlock_Error= %s",eB.Error())
			continue
		}
		if CheckPow4RealBlock(block){
			t.Logf("CheckPow4RealBlock_succeed= %t\n",true)
		}else{
			t.Fatalf("CheckPow4RealBlock_fatal= %t\n",false)
		}
	}


}

func TestCheckCmpLastBlock10(t *testing.T) {

	var lastBlockHeight int

	btcApi := cypherbtc.NewBtcApi(*token)

	chain,err := btcApi.FetchChain()

	if err != nil {
		t.Errorf("FetchChain()_error= %s\n",err.Error())
		return
	}
	lastBlockHeight = chain.Height

	for i := lastBlockHeight ; i > lastBlockHeight - 10 ; i-- {
		fmt.Printf("-------------%d--------------\n",i)
		t.Logf("block_height= %d", i)
		block, eB := btcApi.FetchBlock(i)

		if eB != nil {
			t.Errorf("FetchBlock_Error= %s",eB.Error())
			continue
		}

		if CheckCmp(block){
			t.Logf("CheckPow4RealBlock_succeed= %t\n",true)
		}else{
			t.Fatalf("CheckPow4RealBlock_fatal= %t\n",false)
		}



	}
}