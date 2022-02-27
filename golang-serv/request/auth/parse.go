package auth

type Parse struct {
	Token string `json:"token" validate:"required"`
}

func NewParse() *Parse {
	return &Parse{}
}

func (p *Parse) GetToken() string {
	return p.Token
}

func (p *Parse) SetToken(token string) {
	p.Token = token
}
