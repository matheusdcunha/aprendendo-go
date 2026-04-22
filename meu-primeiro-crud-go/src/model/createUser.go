package model

import (
	"fmt"

	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/config/logger"
	"github.com/matheusdcunha/aprendendo-go/meu-primeiro-crud-go/src/config/resterr"
)

func (ud *userDomain) CreateUser() *resterr.RestErr {
	logger.Info("init CreateUser model", logger.CreateUserJourney)

	ud.EncryptPassword()

	fmt.Println(ud)
	return nil
}
