package model

import "encoding/json"

type Token struct {
	Token string `json:"token"`
}

func GetTokenFromResponse(response string) (*Token, error) {
	t := new(Token)
	err := json.Unmarshal([]byte(response), t)
	return t, err
}
