package cypherbtc

import (
	"gopkg.in/blockcypher/gobcy.v1"
)


type BtcApi struct {
	btcApi	*gobcy.API
}

func NewBtcApi(token string) (api *BtcApi) {

	api = &BtcApi{}
	api.btcApi = &gobcy.API{token,"btc","main"}
	return
}

func (api *BtcApi)FetchChain() (gobcy.Blockchain,error){

	return api.btcApi.GetChain()

}


func (api *BtcApi)FetchBlock(blockNumber int)(block gobcy.Block,err error){


	//"-1" indicate last block
	var number int
	if blockNumber <= -1 {

		chain,e := api.FetchChain()
		if e != nil {
			err = e
			return
		}
		number = chain.Height
	}else{
		number = blockNumber
	}


	b,errBlock := api.btcApi.GetBlock(number,"",nil)
	if errBlock != nil {
		err = errBlock
		return
	}

	block = b
	return
}