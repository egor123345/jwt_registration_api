package regJwt

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	jwt.StandardClaims
	UserId int
}
