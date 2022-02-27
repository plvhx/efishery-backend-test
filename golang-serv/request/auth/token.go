package auth

type Token struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewToken() *Token {
	return &Token{}
}

func (t *Token) GetPhone() string {
	return t.Phone
}

func (t *Token) SetPhone(phone string) {
	t.Phone = phone
}

func (t *Token) GetPassword() string {
	return t.Password
}

func (t *Token) SetPassword(password string) {
	t.Password = password
}
