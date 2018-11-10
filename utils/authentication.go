package utils

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"tc-web-gateway/domains/models"
	user2 "tc-web-gateway/feature/user"
	"tc-web-gateway/utils/errors"
	"testjwt/core/redis"
	"time"
)

type Authentication struct {
	PrivateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
	SymmetricKey string
}

const (
	tokenDuration = 72
	expireOffset = 3600
)

var instance *Authentication = nil

func InitAuthentication() *Authentication{
	if instance == nil {
		instance = &Authentication{
			PrivateKey: getPrivateKey(),
			PublicKey:getPublicKey(),
			SymmetricKey: getSymmetricKey(),
		}
	}
	return instance
}

func (a *Authentication) Authenticate (user *models.User) (bool, error) {

	service := user2.NewUserService()
	userFromDb, err:= service.FindByUsername(user.Username)
	if err != nil{
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(user.Password))
	return user.Username == userFromDb.Username && err == nil, nil
}

func (a *Authentication) GenerateToken(user *models.User) (string, error){
	token:= jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	//TODO change DURATION TO ENV VAR
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(72)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = user.ID
	claims["role"] = user.RoleName
	tokenString,err := token.SignedString([]byte(a.SymmetricKey))
	return tokenString, err
}

func (a *Authentication) CheckTokenValidity(tokenString string) (*jwt.Token, error) {
	var err error
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok:= token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return []byte(a.SymmetricKey), nil
		}
	})

	if err == nil && token.Valid /* && token.NotInBlackList TODO */ {
		return token, nil
	}

	return nil, errors.ErrInvalidToken
}

func (a *Authentication) getTokenRemainingValidity(timestamp interface{})int{
	if validity, ok := timestamp.(float64); ok{
		tm:= time.Unix(int64(validity), 0)
		remainer:= tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}
func (a *Authentication) Logout(tokenString string, token *jwt.Token) error {
	//TODO waiting for regis installed
	redisConn := redis.Connect()
	claims:= token.Claims.(jwt.MapClaims)
	return redisConn.SetValue(tokenString, tokenString, a.getTokenRemainingValidity(claims["exp"]))
}


func (a *Authentication) InvalidateToken(tokenString string, token *jwt.Token) error{
	//TODO
	return nil
}

func getPrivateKey() *rsa.PrivateKey {
	//privateKeyImported, _ := x509.ParsePKCS1PrivateKey([]byte("dummy"))

	//if err != nil {
	//	panic(err)
	//}
	//TODO
	return nil
}

func getPublicKey() *rsa.PublicKey {
	//TODO
	//publicKeyImported, err := x509.ParsePKIXPublicKey([]byte("dummyPublicKey"))

	//if err != nil {
	//	panic(err)
	//}

	//rsaPub, ok := publicKeyImported.(*rsa.PublicKey)
	//
	//if !ok {
	//	panic(err)
	//}

	return nil
}

func getSymmetricKey() string {
	//TODO
	return "secret"
}