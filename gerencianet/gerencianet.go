package gerencianet

type gerencianet struct {
	endpoints
}

func NewGerencianet(configs map[string]interface{}) *gerencianet {
	clientID := configs["client_id"].(string)
	clientSecret := configs["client_secret"].(string)
	CA := configs["CA"].(string)
	Key := configs["Key"].(string)
	sandbox := configs["sandbox"].(bool)
	timeout := configs["timeout"].(int)

	requester := newRequester(clientID, clientSecret,CA, Key, sandbox, timeout)
	gn := gerencianet{}
	gn.requester = *requester
	return &gn
}