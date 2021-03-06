package services

import (
	"errors"
	"strings"

	"github.com/aaraya0/arq-software/ejerciciosGo/ej-auth/domain"
	"github.com/aaraya0/arq-software/ejerciciosGo/ej-auth/utils"
)

const (
	credentialsPath = "credentials.txt"
	tokenPath       = "token.txt"
)

func Login(cred domain.Credentials) (domain.Token, error) {
	bytes, err := utils.ReadFile(credentialsPath)
	if err != nil {
		return domain.Token{}, err
	}

	loggedIn := false
	for _, line := range strings.Split(string(bytes), "\n") {
		components := strings.Split(line, "@")
		user, password := components[0], components[1]
		if user == cred.User && password == cred.Password {
			loggedIn = true
			break
		}
	}

	if !loggedIn {
		return domain.Token{}, errors.New("invalid credentials")
	}

	tokenBytes, err := utils.ReadFile(tokenPath)
	if err != nil {
		return domain.Token{}, err
	}

	return domain.Token{
		Token: string(tokenBytes),
	}, nil
}
