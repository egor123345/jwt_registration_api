package regJwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtServ struct {
	signingKey    string
	hoursOfAction int
}

func NewJwtServ(signingJwtKey string, hoursOfJwtAction int) *JwtServ {
	return &JwtServ{
		signingKey:    signingJwtKey,
		hoursOfAction: hoursOfJwtAction,
	}
}

func (j *JwtServ) GenerateUserToken(userId int) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(j.hoursOfAction) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
	})
	return jwtToken.SignedString([]byte(j.signingKey))
}
