package auth

type Register struct {
	Name  string `json:"name" validate:"required"`
	Phone string `json:"phone" validate:"required"`
	Role  string `json:"role" validate:"required"`
}

func NewRegister() *Register {
	return &Register{}
}

func (r *Register) GetName() string {
	return r.Name
}

func (r *Register) SetName(name string) {
	r.Name = name
}

func (r *Register) GetPhone() string {
	return r.Phone
}

func (r *Register) SetPhone(phone string) {
	r.Phone = phone
}

func (r *Register) GetRole() string {
	return r.Role
}

func (r *Register) SetRole(role string) {
	r.Role = role
}
