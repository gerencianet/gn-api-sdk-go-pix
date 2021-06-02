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
		
		"calendario": map[string]interface{} {
				"expiracao": 3600,
			},
		"devedor": map[string]interface{}{
			
				"cpf": "12345678909",
				"nome": "Francisco da Silva",
			
		},
		"valor": map[string]interface{} {
			
				"original": "00.01",
			
		},
		"chave": "48044e07-e215-417f-b1ad-32ee2d99c2bc",
		"solicitacaoPagador": "Teste.",
		"infoAdicionais": []map[string]interface{} {
			{
				"nome": "Campo 1",
				"valor": "Informação Adicional1 do PSP-Recebedor",
			},
		},
	}

	res, err := gn.CreateImmediateCharge(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}