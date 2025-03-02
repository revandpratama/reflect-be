package helper

import (
	"errors"
	"strconv"
	"sync"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/revandpratama/reflect/auth-service/config"
	"github.com/revandpratama/reflect/auth-service/internal/entities"
)

// var secretKey, _ = paseto.V4SymmetricKeyFromBytes([]byte(config.ENV.SecretKey))

var (
	secretKey paseto.V4SymmetricKey
	once      sync.Once
)

func initSecretKey() {
	once.Do(func() {
		secretKey, _ = paseto.V4SymmetricKeyFromBytes([]byte(config.ENV.SecretKey))
	})
}

func CreateToken(user entities.User) (string, error) {
	initSecretKey()

	token := paseto.NewToken()

	//set rule
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Minute))

	//insert paylaod
	token.SetString("user_id", strconv.Itoa(user.ID))
	token.SetString("role_id", strconv.Itoa(user.RoleID))
	token.SetString("name", user.Name)
	token.SetString("email", user.Email)

	encrypted := token.V4Encrypt(secretKey, nil)

	var err error
	if encrypted == "" {
		return "", errors.New("failed creating token")
	}

	return encrypted, err

}

func VerifyToken(encryptedToken string) (*entities.User, error) {

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
