package cypherbtc

import (
	"encoding/json"
	"flag"
	"fmt"
	"testing"
)

var (
	token *string
)

func init() {
	token = flag.String("token","nil string","blockCypher's token")
	fmt.Printf("token= %s\n",*token)
}

func TestFetchChain(t *testing.T) {


	fmt.Printf("test_token= %s",*token)
	api := NewBtcApi(*token)

	bc,err := api.FetchChain()

	if err != nil {
		t.Errorf("error= %s",err.Error())
		return
	}

	data,err := json.Marshal(bc)
	if err != nil {
		t.Errorf("error= %s",err.Error())
		return
	}

	t.Logf("json_data= %s",string(data))
}


func TestFetchBlock(t *testing.T) {

	api := NewBtcApi(*token)

	block,err := api.FetchBlock(-1)

	if err != nil {
		t.Errorf("error= %s",err.Error())
		return
	}

	bData,eJson := json.Marshal(block)

	if eJson != nil {
		t.Errorf("error= %s",eJson.Error())
		return
	}

	t.Logf("block_json= %s",string(bData))

}