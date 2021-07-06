package models

type Tabler interface {
	TableName() string
}

func (Auth) TableName() string {
	return "auth"
}

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	//Level    string `json:"level"`
}
type AuthEntity interface {
	Cek(username, password string) (Auth, error)
}

type AuthRepository interface {
	Cek(username, password string) (Auth, error)
}
