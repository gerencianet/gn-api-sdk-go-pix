# gn-api-sdk-go

> A go library for integration of your backend with the payment services
provided by [Gerencianet](http://gerencianet.com.br).

[![Build Status](https://travis-ci.org/gerencianet/gn-api-sdk-go.svg)](https://travis-ci.org/gerencianet/gn-api-sdk-go)
[![Coverage Status](https://coveralls.io/repos/github/gerencianet/gn-api-sdk-go/badge.svg?branch=master)](https://coveralls.io/github/gerencianet/gn-api-sdk-go?branch=master)

## Installation

Install with:

```bash

$ go get github.com/gerencianet/gn-api-sdk-go/gerencianet
```
## Tested with
```
go 1.15.5

```
## Basic usage

```go

import (
	"fmt"
	"../../gerencianet"
	"../configs"
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
		"chave": "2c5c7441-a91e-4982-8c25-6105581e18ae",
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
