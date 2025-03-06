package helper

import (
	"strconv"
	"sync"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/revandpratama/reflect/config"
	"github.com/revandpratama/reflect/helper"
	"github.com/revandpratama/reflect/internal/entities"
)

var (
	secretKey paseto.V4SymmetricKey
	once      sync.Once
)

func initSecretKey() {
	once.Do(func() {
		var err error
		secretKey, err = paseto.V4SymmetricKeyFromBytes([]byte(config.ENV.SecretKey))
		if err != nil {
			helper.NewLog().Fatal("Failed to initialize secret key") // Handle this properly in production
		}
	})
}

func VerifyToken(encryptedToken string) (*entities.User, error) {
	initSecretKey()

	parser := paseto.NewParser()
	parser.AddRule(paseto.NotExpired())
	parser.AddRule(paseto.ValidAt(time.Now()))

	parsedToken, err := parser.ParseV4Local(secretKey, encryptedToken, nil)
	if err != nil {
		return nil, err
	}

	user, err := getPayloadFromParsedToken(parsedToken)

	return user, err
}

func getPayloadFromParsedToken(parsedToken *paseto.Token) (*entities.User, error) {
	userIDStr, err := parsedToken.GetString("user_id")
	if err != nil {
		return nil, err
	}
	roleIDStr, err := parsedToken.GetString("role_id")
	if err != nil {
		return nil, err
	}
	name, err := parsedToken.GetString("name")
	if err != nil {
		return nil, err
	}
	email, err := parsedToken.GetString("email")
	if err != nil {
		return nil, err
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return nil, err
	}
	roleID, err := strconv.Atoi(roleIDStr)
	if err != nil {
		return nil, err
	}

	user := entities.User{
		ID:     userID,
		RoleID: roleID,
		Name:   name,
		Email:  email,
	}

	return &user, nil
}
