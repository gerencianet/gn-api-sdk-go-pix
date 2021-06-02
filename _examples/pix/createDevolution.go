package main

import (
	"fmt"
	"../../gerencianet"
	"../configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	
	const e2eid = "E00416968202105211756Rh0iPsaJ1RK"
	const id = "adssshdsjdsjeyccdyddsasdstxid29"

	body := map[string]interface{} {		
		"valor": "0.01",
	}

	res, err := gn.CreateDevolution(e2eid,id,body) // no lugar do 1 coloque o charge_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}