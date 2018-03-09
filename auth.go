package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

var rootpswd = "asdasd"
var key = "secret"
var jwtHeader = `{
	'typ': 'JWT',
	'alg': 'HS256'
  }`
var token Token

func init() {
	token.init([]byte(key), jwtHeader)
}

//Token partly implements JWT for token authorisation
type Token struct {
	key      []byte
	header   tHeader
	playload tPlayload
}
type tHeader struct {
	typ string
	alg string
}
type tPlayload struct {
	exp  string //datetime
	iat  string //datetime
	nbf  string //cms_suer(user_log); nil for no check
	id   string
	Name string  `json:"name"`　
}

func (t *tHeader) put(typ, alg string) {
	if typ == "" {
		typ = "jwt"
	}
	if alg == "" {
		alg = "sha256"
	}
	t.typ = typ
	t.alg = alg
}

func(t *tPlayload) put(exp, iat, nbf, id, name string){

}
func (t *Token) encode(playload map[string]string) (string, error) {
	jsonString, err := json.Marshal(playload)
	if err != nil {
		return "", err
	}
	mac := hmac.New(sha256.New, t.key)

	mac.Write([]byte(t.header))
	mac.Write([]byte("."))
	mac.Write(jsonString)
	output := mac.Sum(nil)
	fmt.Println(string(output))
	return t.header + "." + base64.StdEncoding.EncodeToString(jsonString) + "." + string(output), nil
}

func (t *Token) decode() map[string]string {
	data := make(map[string]string)

	return data

}

func (t *Token) validate() bool {
	return false
}

func (t *Token) init(key []byte, header string) error {
	//base64.StdEncoding.Encode(t.key, key)
	t.header = base64.StdEncoding.EncodeToString([]byte(header))
	t.key = []byte(base64.StdEncoding.EncodeToString(key))
	return nil
}
