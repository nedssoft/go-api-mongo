package requests

import "github.com/nedssoft/go-basic-api/utils"

type UserPayload struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=100,containsany=!@#$%^&*()_+{}[]:<>?~"`
}

type UserUpdatePayload struct {
	Name string `json:"name" binding:"min=3,max=100"`
}


func (u *UserPayload) HashPassword() error {
	hash, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}

func (u *UserUpdatePayload) Validate() error {
	return nil
}

func (u *UserPayload) Validate() error {
	return nil
}

