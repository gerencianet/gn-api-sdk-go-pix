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
		
		"pix": map[string]interface{}{
			    "receberSemChave": false,
				"chaves": map[string]interface{} {		
					"48044e07-e215-417f-b1ad-32ee2d99c2bc":map[string]interface{} {
						"recebimento":map[string]interface{} {
							"txidObrigatorio": false,
							"qrCodeEstatico":map[string]interface{} {
								"recusarTodos": false,

							},

						},

					},

			},
			
		},
		
	}

	res, err := gn.UpdateSettings(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}