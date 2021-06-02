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
	const key = "48044e07-e215-417f-b1ad-32ee2d99c2bc"

	res, err := gn.DeleteKey(key, body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}