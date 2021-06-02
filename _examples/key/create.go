package main

import (
	"fmt"
	"../../gerencianet"
	"../configs"
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