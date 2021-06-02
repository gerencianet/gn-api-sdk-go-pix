package main

import (
	"fmt"
	"../../gerencianet"
	"../configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)



	body := map[string]interface{} {		
		"tipoCob": "cob",
	}

	res, err := gn.CreateLoc(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}