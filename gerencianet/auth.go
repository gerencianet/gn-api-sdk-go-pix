package gerencianet

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"time"
  	"crypto/tls"
)

type auth struct {
	clientID, clientSecret string
	CA, Key				   string
	sandbox                bool
	timeout                int
	netClient              interface {
		Do(req *http.Request) (*http.Response, error)
	}
}

type authResponseBody struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	EspiresAt    string `json:"expire_at"`
	TokenType    string `json:"token_type"`
}


func newAuth(clientID string, clientSecret string,CA string, Key string, sandbox bool, timeout int) *auth {

	
	var cert, erro = tls.LoadX509KeyPair(CA, Key)
	if erro != nil {
		println("Erro ao carregar par de chaves",erro)
	}

	
	var netTransport = &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}
	httpClient := &http.Client{Timeout: time.Second * time.Duration(timeout),Transport: netTransport}
	return &auth{clientID, clientSecret,CA, Key, sandbox, timeout, httpClient}
}



func (auth auth) getAccessToken() (authResponseBody, error) {
	var response authResponseBody
	credentials := auth.clientID + ":" + auth.clientSecret
	authCredentials := b64.StdEncoding.EncodeToString([]byte(credentials))
	postParams := map[string]string{"grant_type": "client_credentials"}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(postParams)

	var gnURL string
	if auth.sandbox {
		gnURL = UrlSandbox
	} else {
		gnURL = UrlProduction
	}
	req, _ := http.NewRequest("POST", gnURL+"/oauth/token", buffer)
	req.Header.Add("Authorization", "Basic "+authCredentials)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("charset", "utf-8")
	req.Header.Add("api-sdk", "go-" + Version)


	res, resErr := auth.netClient.Do(req)

	


	if resErr != nil {
		return response, resErr
	}
	
	if res.StatusCode != http.StatusOK {
		return response, errors.New("Status: " + res.Status + " Could not authenticate. \nPlease make sure you are using correct credentials and if you are using then in the correct environment.")
	}

	json.NewDecoder(res.Body).Decode(&response)

	return response, nil
	
}