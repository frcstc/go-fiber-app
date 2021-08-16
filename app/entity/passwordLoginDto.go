package entity

type PasswordLoginDto struct {
	Mobile string `json:"mobile"`
	Password string `json:"password"`
	MobilePrefix string `json:"mobilePrefix"`
}
