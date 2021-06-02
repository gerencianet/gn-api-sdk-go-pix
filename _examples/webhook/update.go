package main

import (
	"fmt"
	"../../gerencianet"
	"../configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	
	const chave = "937e676f-583b-42df-82cd-ff27599c2bb4"

	body := map[string]interface{} {
		
		"webhookUrl": "https://gnmatheus.igorpedroso.dev:3000/webhook",
	}

	res, err := gn.UpdateWebhook(chave,body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}