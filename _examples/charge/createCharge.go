package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go-pix/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go-pix/_examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	
	const txid = "adssshdsjdsjeyccdyddsasdstxid29"

	body := map[string]interface{} {
		
		"calendario": map[string]interface{} {
				"expiracao": 3600,
			},
		"devedor": map[string]interface{}{
			
				"cnpj": "12345678000195",
				"nome": "Empresa de Serviços SA",
			
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

	res, err := gn.CreateCharge(txid,body) // no lugar do 1 coloque o charge_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}