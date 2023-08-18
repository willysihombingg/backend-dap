package middleware

import (
	"fmt"
	"time"

	"github.com/kataras/jwt"
)

type Claims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func GenerateToken(id int, email string, role string, secKey string, age time.Duration) (string, error) {
	now := time.Now()
	standardClaims := jwt.Claims{
		Expiry:   now.Add(age * time.Hour).Local().Unix(),
		IssuedAt: now.Unix(),
		Issuer:   "Login Activity",
	}

	myClaims := Claims{
		Id:    id,
		Email: email,
		Role:  role,
	}

	token, err := jwt.Sign(jwt.HS256, []byte(secKey), myClaims, standardClaims)

	return string(token), err
}

func VerifyToken(secKey string, tkn string) (*Claims, error) {
	type MyClaims struct {
		ID       int
		Email    string
		Role     string
		Expiry   int64
		IssuedAt int
		Issuer   string
	}

	clm := MyClaims{}
	verifiedToken, err := jwt.Verify(jwt.HS256, []byte(secKey), []byte(tkn))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = verifiedToken.Claims(&clm)

	if err != nil {
		return nil, err
	}

	return &Claims{
		Id:    clm.ID,
		Email: clm.Email,
		Role:  clm.Role,
	}, err
}
