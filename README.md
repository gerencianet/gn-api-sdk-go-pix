![alt text](https://s3.amazonaws.com/gerencianet-pub-prod-1/printscreen/2021/06/29/matheus.rodrigues/68747470733a2f2f6d656469612d657870312e6c6963646e2e636f6d2f646d732f696d6167652f43344431424151483974614e49615a79685f512f636f6d70616e792d6261636b67726f756e645f31303030302f302f313630333132363632333936343f65.jpg)

<h1 align="center">gn-api-sdk-go-pix</h1>

> A go library for integration of your backend with the payment services
provided by [Gerencianet](http://gerencianet.com.br).

[![Build Status](https://travis-ci.org/gerencianet/gn-api-sdk-go.svg)](https://travis-ci.org/gerencianet/gn-api-sdk-go)
[![Coverage Status](https://coveralls.io/repos/github/gerencianet/gn-api-sdk-go/badge.svg?branch=master)](https://coveralls.io/github/gerencianet/gn-api-sdk-go?branch=master)

## Installation

Install with:

```bash

$ go get github.com/gerencianet/gn-api-sdk-go-pix/gerencianet
$ go mod init github.com/gerencianet/gn-api-sdk-go-pix
```
## Tested with
```
go 1.16.5

```
## Certificate conversion
```js
openssl pkcs12 -in path.p12 -out newfile.crt.pem -clcerts -nokeys  //  certificado 
openssl pkcs12 -in path.p12 -out newfile.key.pem -nocerts -nodes //  chave privada
```
## Basic usage

```go

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go-pix/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go-pix/_examples/configs"
)

credentials := map[string]interface{} {
    "client_id": "Your Client_Id",
	"client_secret": "Your Client_Secret",
	"sandbox": true,
	"timeout": 20,
	"CA" : "Path to your public key",
	"Key" : "Path to your private key",
}

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
		"chave": "47e7e515-44d3-42cc-8e1f-ef529b4ba4d1",
		"solicitacaoPagador": "Teste.",
		"infoAdicionais": []map[string]interface{} {
			{
				"nome": "Campo 1",
				"valor": "Informação Adicional1 do PSP-Recebedor",
			},
		},
	}

res, err := gn.CreateImmediateCharge(body)

```

## Examples

You can run the examples inside `_examples` with
`$ go run example.go`:

```bash
$ go run chargecreateImmediateCharge.go
```

Just remember to set the correct credentials inside `examples/configs.go` before running.



## Changelog

[CHANGELOG](CHANGELOG.md)


## License

The library is available as open source under the terms of the [MIT License](LICENSE).
