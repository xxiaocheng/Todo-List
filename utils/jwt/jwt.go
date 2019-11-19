package jwt

import (
	"time"
	"todoList/config"
	"todoList/serializers"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenJwtToken(username string) serializers.JwtResponse {
	expiredTime := time.Now().Add(config.Config.ExpireTime * time.Hour)
	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Set some claims
	jwt_token.Claims = jwt.MapClaims{
		"username": username,
		"exp":      expiredTime.Unix(),
	}
	// Sign and get the complete encoded token as a string
	token, _ := jwt_token.SignedString([]byte(config.Config.Secret))
	return serializers.JwtResponse{
		AccessToken: token,
		ExpireTime:  expiredTime.Unix(),
	}
}

func ParseJwtToken(token string) (*jwt.MapClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.Secret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.MapClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
