package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go-pix/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go-pix/_examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	body := map[string]interface{} {}

	
	res, err := gn.CreateKey(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}