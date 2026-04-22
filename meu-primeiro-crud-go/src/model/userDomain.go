package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/config/resterr"
)

type userDomain struct {
	email    string
	password string
	name     string
	age      int
}

func NewUserDomain(email, password, name string, age int) UserDomainInterface {
	return &userDomain{
		email,
		password,
		name,
		age,
	}
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *resterr.RestErr
	UpdateUser(string) *resterr.RestErr
	FindUser(string) (*userDomain, *resterr.RestErr)
	DeleteUser(string) *resterr.RestErr
}
