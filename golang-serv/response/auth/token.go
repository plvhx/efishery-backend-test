package auth

type Token struct {
	Token string `json:"token"`
}

func NewToken() *Token {
	return &Token{}
}

func (t *Token) GetToken() string {
	return t.Token
}

func (t *Token) SetToken(token string) {
	t.Token = token
}
