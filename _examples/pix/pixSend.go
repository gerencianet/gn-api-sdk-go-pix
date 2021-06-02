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
		"valor": "10.00",
		"pagador": map[string]interface{} {
			"chave": "48044e07-e215-417f-b1ad-32ee2d99c2bc",
			"infoPagador": "Segue o pagamento da conta",
			},
		"favorecido": map[string]interface{} {			
			"chave": "620e2ddb-9746-4e0e-9350-dc6dff699224",
		},
	}

	res, err := gn.PixSend(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}