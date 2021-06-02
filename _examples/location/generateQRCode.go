package main

import (
	"fmt"
	"../../gerencianet"
	"../configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	const id = "423"

	res, err := gn.GenerateQRCode(id) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}