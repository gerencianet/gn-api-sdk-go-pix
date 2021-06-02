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


	res, err := gn.DevolutionList(e2eid,id) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}