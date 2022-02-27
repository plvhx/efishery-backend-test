package auth

type Auth struct {
	Name      string `json:"name" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Role      string `json:"role" validate:"required"`
	Password  string `json:"password,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) GetName() string {
	return a.Name
}

func (a *Auth) SetName(name string) {
	a.Name = name
}

func (a *Auth) GetPhone() string {
	return a.Phone
}

func (a *Auth) SetPhone(phone string) {
	a.Phone = phone
}

func (a *Auth) GetRole() string {
	return a.Role
}

func (a *Auth) SetRole(role string) {
	a.Role = role
}

func (a *Auth) GetPassword() string {
	return a.Password
}

func (a *Auth) SetPassword(password string) {
	a.Password = password
}

func (a *Auth) GetTimestamp() int64 {
	return a.Timestamp
}

func (a *Auth) SetTimestamp(timestamp int64) {
	a.Timestamp = timestamp
}
