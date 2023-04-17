package jwtPkg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const issuer = "golang-backend-microservice"

type authCustomClaims struct {
	Id uint `json:"id"`
}

func GenerateToken(id uint, durationMinute int) string {
	claims := &jwt.MapClaims{
		"id": id,
	}

	if durationMinute > 0 {
		claims = &jwt.MapClaims{
			"id":  id,
			"exp": time.Now().Add(time.Minute * time.Duration(durationMinute)).Unix(),
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		panic(err)
	}
	return t
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid_token_%v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

}

func GetClaims(tokenObj *jwt.Token) (*authCustomClaims, bool) {
	if claims, ok := tokenObj.Claims.(jwt.MapClaims); ok && tokenObj.Valid {
		var authClaim authCustomClaims
		b, err := json.Marshal(claims)
		if err != nil {
			return nil, false
		}
		err = json.Unmarshal(b, &authClaim)
		if err != nil {
			return nil, false
		}
		return &authClaim, true
	} else {
		fmt.Println(reflect.TypeOf(tokenObj.Claims))
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func ValidAndGetClaims(encodedToken string) (*authCustomClaims, bool) {
	tokenObj, err := ValidateToken(encodedToken)
	if err != nil {
		return nil, false
	}
	return GetClaims(tokenObj)
}
