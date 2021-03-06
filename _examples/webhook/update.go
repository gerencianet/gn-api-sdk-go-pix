package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go-pix/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go-pix/_examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	
	const chave = "937e676f-583b-42df-82cd-ff27599c2bb4"

	body := map[string]interface{} {
		
		"webhookUrl": "https://seu_webhook",
	}

	res, err := gn.UpdateWebhook(chave,body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
