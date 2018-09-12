package main

import (
	"fmt"
	"github.com/hoonrepo/goUtils/btcutil"
	"github.com/hoonrepo/goUtils/param"
)



func main(){

	btcutil.Check539753()

	token := param.ConfigParam.BtcConfig.Token

	fmt.Printf("token= %s",token)

}
