package gerencianet


type endpoints struct {
	requester interface {
		request(endpoint string, httpVerb string, requestParams map[string]string, body map[string]interface{}) (string, error)
	}
}




func (endpoints endpoints) CreateImmediateCharge(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/cob", "POST", nil, body)
}


func (endpoints endpoints) CreateCharge(txid string, body map[string]interface{}) (string, error) {
	params := map[string]string{ "txid": (txid) }
	return endpoints.requester.request("/v2/cob/:txid", "PUT", params, body)
}

func (endpoints endpoints) DetailCharge(txid string) (string, error) {
	params := map[string]string{ "txid": (txid) }
	return endpoints.requester.request("/v2/cob/:txid", "GET", params, nil)
}

func (endpoints endpoints) ListCharges(inicio string, fim string) (string, error) {
	params := map[string]string{ 
		"inicio": (inicio),
		"fim": (fim),
	}
	return endpoints.requester.request("/v2/cob?inicio=:inicio&fim=:fim", "GET", params, nil)
}

func (endpoints endpoints) CreateDevolution(e2eid string, id string, body map[string]interface{}) (string, error) {
	params := map[string]string{
		 "e2eid": (e2eid),
		 "id": (id), }
	return endpoints.requester.request("/v2/pix/:e2eid/devolucao/:id", "PUT", params, body)
}

func (endpoints endpoints) DevolutionList(e2eid string, id string) (string, error) {
	params := map[string]string{
		 "e2eid": (e2eid),
		 "id": (id), }
	return endpoints.requester.request("/v2/pix/:e2eid/devolucao/:id", "GET", params, nil)
}

func (endpoints endpoints) PixDetail(e2eid string) (string, error) {
	params := map[string]string{ "e2eid": (e2eid) }
	return endpoints.requester.request("/v2/pix/:e2eid", "GET", params, nil)
}

func (endpoints endpoints) PixSend(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/pix", "POST", nil, body)
}

func (endpoints endpoints) CreateLoc(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/loc", "POST", nil, body)
}

func (endpoints endpoints) DeleteTxid(id string, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": (id) }
	return endpoints.requester.request("/v2/loc/:id/txid", "DELETE", params, body)
}

func (endpoints endpoints) GetLoc(id string) (string, error) {
	params := map[string]string{ "id": (id) }
	return endpoints.requester.request("/v2/loc/:id", "GET", params, nil)
}

func (endpoints endpoints) ListLoc(inicio string, fim string) (string, error) {
	params := map[string]string{ 
		"inicio": (inicio),
		"fim": (fim),
	}
	return endpoints.requester.request("/v2/loc?inicio=:inicio&fim=:fim", "GET", params, nil)
}

func (endpoints endpoints) GenerateQRCode(id string) (string, error) {
	params := map[string]string{ "id": (id) }
	return endpoints.requester.request("/v2/loc/:id/qrcode", "GET", params, nil)
}

func (endpoints endpoints) Balance(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/gn/saldo", "GET", nil, body)
}
func (endpoints endpoints) ListSettings(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/gn/config", "GET", nil, body)
}

func (endpoints endpoints) UpdateSettings(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/gn/config", "PUT", nil, body)
}
func (endpoints endpoints) CreateKey(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/gn/evp", "POST", nil, body)
}

func (endpoints endpoints) ListKey(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/v2/gn/evp", "GET", nil, body)
}

func (endpoints endpoints) DeleteKey(key string, body map[string]interface{}) (string, error) {
	params := map[string]string{ "key": (key) }
	return endpoints.requester.request("/v2/gn/evp/:key", "DELETE", params, body)
}

func (endpoints endpoints) UpdateWebhook(chave string, body map[string]interface{}) (string, error) {
	params := map[string]string{ "chave": (chave) }
	return endpoints.requester.request("/v2/webhook/:chave", "PUT", params, body)
}

func (endpoints endpoints) DeleteWebhook(chave string, body map[string]interface{}) (string, error) {
	params := map[string]string{ "chave": (chave) }
	return endpoints.requester.request("/v2/webhook/:chave", "DELETE", params, body)
}

func (endpoints endpoints) GetWebhook(chave string) (string, error) {
	params := map[string]string{ "chave": (chave) }
	return endpoints.requester.request("/v2/webhook/:chave", "GET", params, nil)
}

func (endpoints endpoints) ListWebhooks(inicio string, fim string) (string, error) {
	params := map[string]string{ 
		"inicio": (inicio),
		"fim": (fim),
	}
	return endpoints.requester.request("/v2/webhook?inicio=:inicio&fim=:fim", "GET", params, nil)
}